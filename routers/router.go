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
		beego.NSRouter("/register", &controllers.LoginController{},"get,post:Register"),
		beego.NSRouter("/login", &controllers.LoginController{},"get,post:Login"),
		beego.NSRouter("/logout", &controllers.LoginController{},"get,post:Logout"),
		beego.NSRouter("/upload/Img", &controllers.AdminController{},"post:UploadImg"),
		beego.NSRouter("/upload/Article", &controllers.ArticleController{},"post:UploadArticle"),
	)
	Article := beego.NewNamespace("/admin/article",
		beego.NSRouter("/AddArticle", &controllers.ArticleController{},"get,post:AddArticle"),
		beego.NSRouter("/ArticleList", &controllers.ArticleController{},"get,post:ArticleList"),
		beego.NSRouter("/ModArticle", &controllers.ArticleController{},"get,post:ModArticle"),
		beego.NSRouter("/DelArticle", &controllers.ArticleController{},"get:DelArticle"),

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
	BlogType := beego.NewNamespace("/admin/blogtype" ,
		beego.NSRouter("/List",&controllers.BlogTypeController{}, "get:List"),
		beego.NSRouter("/Add",&controllers.BlogTypeController{}, "get,post:Add"),
		beego.NSRouter("/Mod/:id([0-9]+)",&controllers.BlogTypeController{}, "get,post:Mod"),
		beego.NSRouter("/Del/:id([0-9]+)",&controllers.BlogTypeController{}, "get:Del"),
	)
	beego.AddNamespace(Admins)
	beego.AddNamespace(Rbac)
	beego.AddNamespace(BlogType,Article)
}
