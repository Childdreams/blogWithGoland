package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"html/template"
)

type BaseController struct {
	beego.Controller
}


func (c *BaseController)Prepare()  {
	fmt.Println("后台中间件")
}

func (c *BaseController)Render() error{
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.Data["v"]="2211"
	fmt.Println("渲染模板时调用")
	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}
	rb, err := c.RenderBytes()
	if err != nil{

	}
	c.Data["info"] = "TEST"
	fmt.Println(c.Data)
	return c.Ctx.Output.Body(rb)
}