package models

type RoleUser struct {
	UserId int 		`orm:"pk"`
	RoleId int
}


