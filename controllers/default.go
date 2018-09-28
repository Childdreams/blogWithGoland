package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["SelfIntroduction"] = "<b>鲍峰</b>，一个什么都想学，却缺少正式项目的猴子，分享一些技术心得。"
	c.TplName = "index.html"
}

func (c *MainController) About()  {
	c.TplName = "about.html"
}

func (c *MainController) Share()  {
	c.TplName = "share.html"
}

func (c *MainController) Gbook()  {
	c.TplName = "gbook.html"
}

func (c *MainController) Infopic()  {
	c.TplName = "infopic.html"
}

func (c *MainController) List()  {
	c.TplName = "list.html"
}

func (c *MainController) Info()  {
	c.TplName = "info.html"
}


func (c *MainController) Index() {
	c.Data["Website"] = "test"
	c.Data["Email"] = "TEST"
	c.TplName = "index.tpl"
}
