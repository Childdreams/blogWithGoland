package models

import "time"

type User struct {
	Id         int 			`orm:"pk"`
	Username   string
	Password   string
	Email      string 		`orm:"unique"`
	LastTime   time.Time	`orm:"default()"`
	LastIp     string		`orm:"default()"`
	State      int8			`orm:"default()"`
	Created    time.Time	`orm:"default()"`
	Updated    time.Time	`orm:"default()"`
}

