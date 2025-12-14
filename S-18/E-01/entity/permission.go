package entity

type Permission struct {
	ID    uint
	Title PermissionTitle
}

type PermissionTitle string

const (
	UserListPermissions  = PermissionTitle("user-list")
	UserDeletePermission = PermissionTitle("user-delete")
)
