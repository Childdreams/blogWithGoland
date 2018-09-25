package utils

import "time"

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
