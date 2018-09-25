package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

var (
	O orm.Ormer
	)
func init()  {
	orm.RegisterModel(
		new(Users),new(Permissions),new(Roles),new(PermissionRole),new(RoleUser),new(Blogtype),new(Blog),new(BbBlogContent),
		)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/blog?charset=utf8&loc=Asia%2FShanghai")
	O =orm.NewOrm()
	configInfo := beego.AppConfig.String("MYSQL::dbname")
	O.Using(configInfo)
}