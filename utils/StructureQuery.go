package utils

import (
	"time"
	"myproject/models"
)

type RolesSQ struct {
	Id int
	Name string
	Display_name string
	Description string
	Created_at time.Time
	Updated_at time.Time
	Dname string
}

type UserSQL struct {
	Id int
	Name string
	Email string
	RememberToken string
	Created_at time.Time
	Updated_at time.Time
	Dname string
}

type BlogContent struct {
	models.Blog
	Content string
}

type SessionUserInfo struct {
	Router []PermissionsCount
	UserInfo Userinfo
}

type Userinfo struct {
	Email string
}