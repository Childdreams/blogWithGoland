package models

type BbBlogContent struct {
	Id int `orm:"pk"`
	Content string `orm:"null;type(text)"`
} 
