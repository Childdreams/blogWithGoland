package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"html/template"
	"github.com/astaxie/beego/orm"
	"myproject/models"
	"myproject/utils"
)

type BaseController struct {
	beego.Controller
}


func (c *BaseController)Prepare()  {
	fmt.Println("后台中间件")
}

func (c *BaseController)Render() error{
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken()


	o := orm.NewOrm()
	cond := orm.NewCondition()
	Render_permissions := cond.And("is_menu" , 1)
	qs := o.QueryTable("permissions").SetCond(Render_permissions)
	var permissions []*models.Permissions
	_,err := qs.All(&permissions)

	if err != nil {

	}
	permissionTree := utils.GetAllPerm(permissions,0,0)
	c.Data["Render_sidebar"] = permissionTree

	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}
	rb, err := c.RenderBytes()
	if err != nil{

	}

	return c.Ctx.Output.Body(rb)
}