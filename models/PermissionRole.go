package models

type PermissionRole struct {
	PermissionId int `orm:"pk"`
	RoleId int64
}
