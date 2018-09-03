package models

import "time"

type Permissions struct {
	Id int 					`orm:"pk"`
	Name string				`orm:"unique"`
	Display_name string		`orm:"null"`
	Description string		`orm:"null"`
	Parent_id int			`orm:"default(0)"`
	Is_menu int			`orm:"default(0)"`
	Created_at time.Time	`orm:"default(time)"`
	Updated_at time.Time	`orm:"default(time)"`

}