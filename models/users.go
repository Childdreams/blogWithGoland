package models

import "time"

type Users struct {
	Id         		int 		`orm:"pk"`
	Name   			string
	Email      		string 		`orm:"unique"`
	Password   		string		`orm:"default()"`
	RememberToken    string
	CreatedAt    	time.Time	`orm:"default(time)"`
	UpdatedAt  		time.Time	`orm:"default(time)"`
}

