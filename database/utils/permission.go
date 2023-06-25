package utils

// - [U] User
// - [I] Instance
// - [S] System
// - [C] control panel
// - [AL]  All
type Permission = Flag

const (
	ADMIN Permission = 1 << iota // [AL]

	CHANGE_NICKNAME          // [U] change self nick
	ADMIN_CHANGE_DESCRIPTION // [U] change user admin description
	MANAGE_USER              // [U] manage user
	MANAGE_USER_NICKNAME     // [U] change user nick
	MANAGE_USER_INSTANCE     // [U] change user instance
	READ_INSTANCE_STATUS     // [I] read instance status
	CHANGE_INSTANCE_STATUS   // [I] change instance status
	READ_CONSOLE             // [I, C] read instance console
	WRITE_CONSOLE            // [I, S, C] write instance console or system console
	READ_FILE                // [I, S] read file
	WRITE_FILE               // [I, S] write file
	VIEW_AUDIT_LOG           // [I, C] view log

	NONE Permission = 0 // [AL]
)
