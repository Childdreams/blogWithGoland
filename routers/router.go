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
		beego.NSRouter("/upload/Img", &controllers.AdminController{},"post:UploadImg"),

		beego.NSInclude(
			&controllers.AdminController{},
		),
	)
	Rbac := beego.NewNamespace("/admin/rbac",
		//Permissions
		beego.NSRouter("/addpermissions",&controllers.RbacController{},"get,post:AddPermissions"),
		beego.NSRouter("/permissionsLists",&controllers.RbacController{},"get,post:PermissionsLists"),
		beego.NSRouter("/delPsermission",&controllers.RbacController{},"get,post:DelPermission"),
		beego.NSRouter("/ModPermission",&controllers.RbacController{},"get,post:ModPermission"),
		//Role
		beego.NSRouter("/rolelist",&controllers.RbacController{},"get:RoleList"),
		beego.NSRouter("/addrole",&controllers.RbacController{},"get,post:AddRole"),
		beego.NSRouter("/molrole",&controllers.RbacController{},"get,post:ModRole"),
		beego.NSRouter("/delrole",&controllers.RbacController{},"get:DelRole"),
		//User
		beego.NSRouter("/userList",&controllers.RbacController{},"get:UserList"),
		beego.NSRouter("/AddUser",&controllers.RbacController{},"get,post:AddUser"),
		beego.NSRouter("/ModUser",&controllers.RbacController{},"get,post:ModUser"),
		beego.NSRouter("/DelUser",&controllers.RbacController{},"get:DelUser"),
		)
	beego.AddNamespace(Admins)
	beego.AddNamespace(Rbac)
}
