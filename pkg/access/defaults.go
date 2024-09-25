package access

var defaultUsers = map[string]User{
	"alice": {
		Roles: []string{"reader", "writer"},
	},
	"bob": {
		Roles: []string{"writer"},
	},
	"charli": {
		Roles: []string{"reader"},
	},
}

var defaultActions = map[string]Action{
	"get": {
		PermittedRoles: []string{"reader"},
	},
	"set": {
		PermittedRoles: []string{"writer"},
	},
}
