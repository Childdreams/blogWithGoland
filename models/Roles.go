package models

import "time"

type Roles struct {
	Id int 					`orm:"pk"`
	Name string				`orm:"unique"`
	Display_name string		`orm:"null"`
	Description string		`orm:"null"`
	Created_at time.Time	`orm:"default(time)"`
	Updated_at time.Time	`orm:"default(time)"`
}


