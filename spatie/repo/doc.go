package repo

// Usage with spatie facade:

// import "server/src/support/spatie"

// // create roles & permissions (variadic)
// spatie.CreateRole("admin", "editor")
// spatie.CreatePermission("manage_users", "export_data")

// // attach permissions to a role (variadic)
// spatie.AssignPermissionToRole("admin", "manage_users", "export_data")

// // give user one or more roles
// spatie.GiveRoleToUser("user-uuid", "admin", "editor")

// // give user one or more direct permissions
// spatie.GivePermissionToUser("user-uuid", "export_data", "manage_users")

// // revoke one or more permissions from user
// spatie.RevokePermissionFromUser("user-uuid", "manage_users", "export_data")
// ...existing code...
