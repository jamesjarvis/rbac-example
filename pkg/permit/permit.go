package permit

import (
	"fmt"

	"github.com/permitio/permit-golang/pkg/enforcement"
)

// PermitClient is the permit.io client interface.
type PermitClient interface {
	Check(user enforcement.User, action enforcement.Action, resource enforcement.Resource) (bool, error)
}

// Permit handles access control for each user and each action.
// This implementation defers logic to permit.io
type Permit struct {
	permitClient PermitClient
}

// New returns a new instance of the access client, using the given permitClient.
func New(
	permitClient PermitClient,
) *Permit {
	return &Permit{
		permitClient: permitClient,
	}
}

// Check returns true if the user is allowed to perform the action, false otherwise.
func (p *Permit) Check(userID string, actionID string) bool {
	user := enforcement.UserBuilder(userID).Build()
	action := enforcement.Action(actionID)
	resource := enforcement.ResourceBuilder("map").Build()

	allowed, err := p.permitClient.Check(user, action, resource)
	if err != nil {
		fmt.Printf("permit error: %s\n", err.Error())
		return false
	}

	return allowed
}
