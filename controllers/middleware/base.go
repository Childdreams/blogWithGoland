package middleware

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "myproject/models"
	"html/template"
	"myproject/utils"
	"fmt"
)

type BaseController struct {
	beego.Controller
}


func (c *BaseController)Prepare()  {
	permissionTree := c.GetSession("UserInfo")
	if permissionTree != nil {
		vv := permissionTree.(utils.SessionUserInfo)
		flag := false
		for _, routers := range vv.Router {
			fmt.Println(routers.Name)
			if c.Data["RouterPattern"] == routers.Name {
				flag = true
				break
			}
		}
		if flag == false {
			c.Redirect("404", 404)
		}
	}else {
		c.Redirect("/admin/login" , 302)
	}
}

func (c *BaseController)Render() error{
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken()
	permissionTree := c.GetSession("UserInfo")
	if permissionTree != nil {
		vv := permissionTree.(utils.SessionUserInfo)
		c.Data["Render_sidebar"] = vv.Router
	}
	if permissionTree == nil  {
		c.Redirect("/admin/login" , 302)
	}
	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}

	rb, err := c.RenderBytes()

	if err != nil{

	}

	return c.Ctx.Output.Body(rb)
}