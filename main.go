package main

import (
	_ "myproject/models"
	_ "myproject/routers"
	"github.com/astaxie/beego"
	"myproject/controllers"
	"github.com/astaxie/beego/orm"
	"myproject/utils"
)


func main() {
	beego.ErrorController(&controllers.AdminController{})
	if beego.AppConfig.String("debug") == "true" {
		orm.Debug = true
		beego.BConfig.EnableErrorsShow = true
		beego.BConfig.Log.AccessLogs = true
	}
	beego.SetStaticPath("/static","static")
	beego.AddFuncMap("hi",utils.ForAdd)
	beego.AddFuncMap("In_array",utils.In_array)

	beego.AddTemplateExt("html")
	beego.Run()
}

