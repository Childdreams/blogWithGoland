package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"myproject/models"
	"crypto/md5"
	"encoding/hex"
	"myproject/utils"
	"myproject/controllers/middleware"
)

type LoginController struct {
	middleware.LoginBaseMiddleware
}


func (c *LoginController) Login()  {

	if c.Ctx.Request.Method == "GET" {
		if v := c.GetSession("email") ;v != nil{
			c.Redirect("/admin",302)
			c.StopRun()
		}
		c.TplName = "admin/login.html"
	}else {
		Email := c.GetString("email")
		Passwored := utils.EnMd5(c.GetString("password"))
		o := orm.NewOrm()
		qs := o.QueryTable("users")
		var userRes models.Users
		err := qs.Filter("email",Email).Filter("password",Passwored).One(&userRes)
		flask := beego.NewFlash()
		if err != nil {
			flask.Error("Email is not exsit!")
			flask.Store(&c.Controller)
			c.Ctx.Redirect(302,"/admin/login")
			c.StopRun()
		}
		qb , _ := orm.NewQueryBuilder("mysql")
		qb.Select("permissions.*").
			From("roles").
			LeftJoin("permission_role").
			On("permission_role.role_id = roles.id").
			LeftJoin("permissions").
			On("permissions.id = permission_role.permission_id").
			LeftJoin("role_user").
			On("role_user.role_id = roles.id").
			Where("role_user.user_id = ?").
			OrderBy("serialnum").
			Desc()
		sql := qb.String()
		var permissionListSession []*models.Permissions
		o.Raw(sql,userRes.Id).QueryRows(&permissionListSession)
		fmt.Println(sql)
		routers := utils.GetAllPerm(permissionListSession , 0 , 0 )
		//fmt.Println(sql)
		flask.Error("Email is not exsit!")
		flask.Store(&c.Controller)
		var user utils.Userinfo
		user.Email =Email
		UserInfo := utils.SessionUserInfo{routers,user}
		c.SetSession("UserInfo",UserInfo)
		c.Redirect("/admin/",302)
	}

}

func (c *LoginController)Logout() {
	flash := beego.NewFlash()
	if Email := c.GetSession("email") ; Email != "" {
		c.DelSession("email")
		flash.Notice("Logout successful")
		c.Redirect("/admin/login",302)
	}
	fmt.Println("this is add flash")
	flash.Error("Not logged in")
	flash.Store(&c.Controller)
	c.Redirect("/admin/login",302)
}


func (c *LoginController) Register()  {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "admin/register.html"
	}else {
		FirstName := c.GetString("firstName")
		LastName := c.GetString("lastName")
		Email := c.GetString("email")
		Password := c.GetString("password")
		ConfirmPassword := c.GetString("confirmPassword")
		flash:=beego.NewFlash()

		if Password != ConfirmPassword {
			flash.Error("The passwords do not match")
			flash.Store(&c.Controller)
			c.Ctx.Redirect(303,"register")
			c.StopRun()
		}
		o := orm.NewOrm()
		user := new(models.Users)
		user.Name = FirstName + LastName
		h := md5.New()
		h.Write([]byte(Password))
		cipherStr := h.Sum(nil)
		user.Password = hex.EncodeToString(cipherStr)
		user.Email = Email
		fmt.Println("User execute Insert ")
		_ ,err := o.Insert(user)
		c.SetSession("email",Email)
		if err != nil {
			flash.Error("Mail to repeat")
			flash.Store(&c.Controller)
			c.Ctx.Redirect(303,"register")
			c.StopRun()
		}
		c.Redirect("/admin",303)
	}
}
