package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	_ "myproject/models"
	"myproject/models"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"myproject/controllers/middleware"
)

type BlogTypeController struct {
	middleware.BaseController
}

func (c *BlogTypeController)List()  {
	var blogtype []*models.Blogtype
	orm.NewOrm().QueryTable("Blogtype").All(&blogtype)
	fmt.Println(blogtype)
	c.Data["data"] = blogtype
	c.TplName = "admin/BlogType/List.html"



}

func (c *BlogTypeController)Add()  {
	if c.Ctx.Request.Method == "GET"{
		c.TplName = "admin/BlogType/Add.html"
	}else {
		flash := beego.NewFlash()
		Name := c.GetString("name")
		SerialNum := c.GetString("serialnum")
		var blogtype models.Blogtype
		blogtype.Blogname = Name
		blogtype.Created_at = time.Now()
		blogtype.Updated_at = time.Now()
		blogtype.Serialnum ,_ = strconv.Atoi(SerialNum)
		Id ,_ := models.O.Insert(&blogtype)
		if Id < 1{
			flash.Error("Insert ERROR !!!")
		}
		flash.Notice("SUCCESS")
		flash.Store(&c.Controller)
		c.Redirect("Add",302)
	}
}

func (c *BlogTypeController)Mod()  {
	if c.Ctx.Request.Method == "GET" {
		Id := c.GetString("id")
		var blog models.Blogtype
		orm.NewOrm().QueryTable("blogtype").Filter("id",Id).One(&blog)
		c.Data["data"] = blog
		c.TplName = "admin/BlogType/Mod.html"
	}else {
		flash := beego.NewFlash()
		Id := c.GetString("id")
		Name := c.GetString("name")
		Serialnum := c.GetString("serialnum")
		num ,_ := orm.NewOrm().QueryTable("blogtype").Filter("id",Id).Update(orm.Params{"blogname":Name,"serialnum":Serialnum})
		if num < 1{
			flash.Error("Insert ERROR !!!")
		}
		flash.Notice("SUCCESS")
		flash.Store(&c.Controller)
		c.Redirect("../List",302)
	}
}

func (c *BlogTypeController)Del()  {
	flash := beego.NewFlash()
	Id := c.Ctx.Input.Param(":id")
	num ,_ := orm.NewOrm().QueryTable("blogtype").Filter("id",Id).Delete()
	if num < 1{
		flash.Error("Insert ERROR !!!")
	}
	flash.Notice("SUCCESS")
	flash.Store(&c.Controller)
	c.Redirect("../List",302)
}

