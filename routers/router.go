package routers

import (
	"myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/getmealtag", &controllers.MainController{},"get:Index")
    beego.Router("/", &controllers.MainController{})
    beego.Router("/about", &controllers.MainController{} , "get:About")
    beego.Router("/share", &controllers.MainController{} , "get:Share")
    beego.Router("/info", &controllers.MainController{} , "get:Info")
    beego.Router("/gbook", &controllers.MainController{} , "get:Gbook")
    beego.Router("/infopic", &controllers.MainController{} , "get:Infopic")
    beego.Router("/list", &controllers.MainController{} , "get:List")

	Admins :=beego.NewNamespace("/admin",
		beego.NSRouter("/", &controllers.AdminController{},"get:Index"),
		beego.NSRouter("/register", &controllers.AdminController{},"get,post:Register"),
		beego.NSRouter("/login", &controllers.AdminController{},"get,post:Login"),
		beego.NSRouter("/logout", &controllers.AdminController{},"get,post:Logout"),
		beego.NSRouter("/AddArticle", &controllers.AdminController{},"get,post:AddArticle"),
		beego.NSInclude(
			&controllers.AdminController{},
		),
	)
	beego.AddNamespace(Admins)
}
