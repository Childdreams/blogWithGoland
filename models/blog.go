package models

import "time"

type Blog struct {
	Id int 								`orm:"pk"`
	Title string
	Keywords string
	Catalog_id int
	Blog_content_id int
	Blog_content_last_update time.Time
	Type int
	Status int
	Views string
	Created time.Time
}
