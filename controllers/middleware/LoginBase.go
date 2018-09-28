package middleware

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "myproject/models"
	"html/template"
)

type LoginBaseMiddleware struct {
	beego.Controller
}


func (c *LoginBaseMiddleware)Prepare()  {
}

func (c *LoginBaseMiddleware)Render() error{
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken()
	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}

	rb, err := c.RenderBytes()

	if err != nil{

	}

	return c.Ctx.Output.Body(rb)
}