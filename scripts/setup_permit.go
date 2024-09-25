package main

import (
	"context"
	"flag"

	"github.com/jamesjarvis/rbac-example/pkg/service"
	"github.com/permitio/permit-golang/pkg/config"
	"github.com/permitio/permit-golang/pkg/models"
	permitio "github.com/permitio/permit-golang/pkg/permit"
)

var serviceClient *service.Service

func main() {
	// Initialise arguments.
	apiKey := flag.String("permit_api_key", "", "API key for Permit authentication")
	flag.Parse()
	if *apiKey == "" {
		panic("api key required")
	}

	ctx := context.Background()
	permitClient := permitio.New(config.NewConfigBuilder(*apiKey).WithPdpUrl("https://cloudpdp.api.permit.io").Build())

	// Set up actions
	actionGet := models.NewActionBlockEditable()
	actionGet.SetName("get")
	actionSet := models.NewActionBlockEditable()
	actionSet.SetName("set")
	resourceMap := models.NewResourceCreate("map", "Map", map[string]models.ActionBlockEditable{"get": *actionGet, "set": *actionSet})
	resourceMap.SetName("map")
	_, err := permitClient.Api.Resources.Create(ctx, *resourceMap)
	if err != nil {
		panic(err)
	}

	// Set up roles
	roleReader := models.NewRoleCreate("reader", "Reader")
	roleReader.SetPermissions([]string{"map:get"})
	_, err = permitClient.Api.Roles.Create(ctx, *roleReader)
	if err != nil {
		panic(err)
	}
	roleWriter := models.NewRoleCreate("writer", "Writer")
	roleWriter.SetPermissions([]string{"map:set"})
	_, err = permitClient.Api.Roles.Create(ctx, *roleWriter)
	if err != nil {
		panic(err)
	}

	// Set up users
	alice := models.NewUserCreate("alice")
	_, err = permitClient.Api.Users.Create(ctx, *alice)
	if err != nil {
		panic(err)
	}
	bob := models.NewUserCreate("bob")
	_, err = permitClient.Api.Users.Create(ctx, *bob)
	if err != nil {
		panic(err)
	}
	charli := models.NewUserCreate("charli")
	_, err = permitClient.Api.Users.Create(ctx, *charli)
	if err != nil {
		panic(err)
	}

	// Set up roles on users
	_, err = permitClient.Api.Users.AssignRole(ctx, "alice", "reader", "default")
	if err != nil {
		panic(err)
	}
	_, err = permitClient.Api.Users.AssignRole(ctx, "alice", "writer", "default")
	if err != nil {
		panic(err)
	}
	_, err = permitClient.Api.Users.AssignRole(ctx, "bob", "writer", "default")
	if err != nil {
		panic(err)
	}
	_, err = permitClient.Api.Users.AssignRole(ctx, "charli", "reader", "default")
	if err != nil {
		panic(err)
	}
}
