package main

import (
	_ "myproject/models"
	_ "myproject/routers"
	"github.com/astaxie/beego"
	"myproject/controllers"
	"github.com/astaxie/beego/orm"
)

func main() {
	beego.ErrorController(&controllers.AdminController{})
	if beego.AppConfig.String("debug") == "true" {
		orm.Debug = true
	}
	beego.AddTemplateExt("html")
	beego.Run()
}

