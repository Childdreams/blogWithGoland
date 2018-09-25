package models

import "time"

type Blog struct {
	Id int 							`orm:"pk"`
	Title string					`orm:"null"`
	Keywords string					`orm:"null"`
	Catalog_id int					`orm:"null"`
	Blog_content_id int				`orm:"null"`
	Blog_content_last_update int	`orm:"null"`
	Type int						`orm:"null"`
	Status int						`orm:"null"`
	Views int						`orm:"null"`
	Created time.Time
}