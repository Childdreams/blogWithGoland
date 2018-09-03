package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

func init()  {
	orm.RegisterModel(new(User),new(Permissions),new(Roles),new(PermissionRole))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/blog?charset=utf8&loc=Asia%2FShanghai")
	o :=orm.NewOrm()
	configInfo := beego.AppConfig.String("MYSQL::dbname")
	o.Using(configInfo)
}