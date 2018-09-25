package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Blogtype struct {
	Id int 					`orm:"pk"`
	Blogname string			`orm:"unique"`
	Serialnum int			`orm:"default(0);size(5)"`
	Created_at time.Time
	Updated_at time.Time
}

func (b *Blogtype)InsertDB(blogtype Blogtype) (int ,error) {
	id , err := orm.NewOrm().Insert(blogtype)
	return int(id) , err
}

