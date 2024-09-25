package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// Global in-memory map to store key-value pairs
var store = make(map[string]string)
var mu sync.RWMutex // Mutex to handle concurrent access

func main() {
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
	vars := mux.Vars(r)
	key := vars["key"]

	mu.RLock() // Lock for reading
	value, ok := store[key]
	mu.RUnlock()

	if ok {
		http.StatusUnauthorized
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, value)
	} else {
		http.Error(w, "Key not found", http.StatusNotFound)
	}
}

// POST /v1/map/{key} - Handler to set a value by key
func setValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	// Read the request body to get the value
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	value := string(body)

	mu.Lock() // Lock for writing
	store[key] = value
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Value set successfully")
}
