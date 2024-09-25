package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jamesjarvis/rbac-example/pkg/access"
	"github.com/jamesjarvis/rbac-example/pkg/permit"
	"github.com/jamesjarvis/rbac-example/pkg/service"
	"github.com/jamesjarvis/rbac-example/pkg/storage"
	"github.com/permitio/permit-golang/pkg/config"
	permitio "github.com/permitio/permit-golang/pkg/permit"
)

var serviceClient *service.Service

func main() {
	// Initialise arguments.
	apiKey := flag.String("permit_api_key", "", "API key for Permit authentication")
	flag.Parse()

	// Initialise Service.
	storageClient := storage.New()
	var accessClient service.AccessClient
	if *apiKey != "" {
		permitClient := permitio.New(config.NewConfigBuilder(*apiKey).WithPdpUrl("https://cloudpdp.api.permit.io").Build())
		accessClient = permit.New(permitClient)
	} else {
		accessClient = access.New()
	}

	// Instantiate global service
	serviceClient = service.New(storageClient, accessClient)

	// Initialize a new router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/v1/map/{key}", getValue).Methods("GET")
	r.HandleFunc("/v1/map/{key}", setValue).Methods("POST")

	// Start the HTTP server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// GET /v1/map/{key} - Handler to get a value by key
func getValue(w http.ResponseWriter, r *http.Request) {
	key, err := getKey(r)
	if err != nil {
		handleError(w, err)
		return
	}

	userID, err := getUserID(r)
	if err != nil {
		handleError(w, err)
		return
	}

	value, err := serviceClient.Get(userID, key)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, value)
}

// POST /v1/map/{key} - Handler to set a value by key
func setValue(w http.ResponseWriter, r *http.Request) {
	key, err := getKey(r)
	if err != nil {
		handleError(w, err)
		return
	}

	userID, err := getUserID(r)
	if err != nil {
		handleError(w, err)
		return
	}

	// Read the request body to get the value
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	value := string(body)

	err = serviceClient.Set(userID, key, value)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, value)
}

// getKey extracts the key from the request url parameters.
func getKey(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	key, ok := vars["key"]
	if !ok {
		return "", errors.New("key not found")
	}
	return key, nil
}

// getUserID extracts the user value from the request headers.
func getUserID(r *http.Request) (string, error) {
	userHeader, ok := r.Header["User"]
	if !ok || len(userHeader) != 1 {
		return "", errors.New("User not found in headers")
	}
	userID := userHeader[0]
	return userID, nil
}

// handleError sets http status codes for the given error.
func handleError(w http.ResponseWriter, err error) {
	if errors.Is(err, service.Error_UNAUTHORISED) {
		http.Error(w, "UNAUTHORISED", http.StatusUnauthorized)
		return
	}
	if errors.Is(err, service.Error_NOTFOUND) {
		http.Error(w, "NOT_FOUND", http.StatusNotFound)
		return
	}
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}
