package access

import "slices"

// User is a user object, containing the roles this user is a member of.
type User struct {
	Roles []string
}

// Action is an action that can be performed, along with the roles of users allowed to perform this action.
type Action struct {
	PermittedRoles []string
}

// Access handles access control for each user and each action.
type Access struct {
	users   map[string]User
	actions map[string]Action
}

// New returns a new instance of the access client.
// For the purposes of this example application, we are going to use hardcoded
// defaults for the users and actions, defined in defaults.go
func New() *Access {
	return &Access{
		users:   defaultUsers,
		actions: defaultActions,
	}
}

// Check returns true if the user is allowed to perform the action, false otherwise.
func (a *Access) Check(userID string, actionID string) bool {
	user, ok := a.users[userID]
	if !ok {
		return false
	}

	action, ok := a.actions[actionID]
	if !ok {
		return false
	}

	for _, role := range user.Roles {
		if slices.Contains(action.PermittedRoles, role) {
			return true
		}
	}
	return false
}
