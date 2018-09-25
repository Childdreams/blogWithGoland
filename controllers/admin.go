package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	_ "myproject/models"
	"github.com/astaxie/beego/orm"
	"myproject/models"
	"time"
	"fmt"
	"github.com/astaxie/beego"
	"crypto/md5"
	"encoding/hex"
	"myproject/utils"
	"os"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) URLMapping() {
	c.Mapping("forgot-password", c.ForgotPassword)
	c.Mapping("error", c.Error404)
	c.Mapping("blank", c.Blank)
	c.Mapping("charts", c.Charts)
	c.Mapping("tables", c.Tables)
	c.Mapping("article", c.Article)
}


func (c *AdminController) Index()  {
	v := c.GetSession("email")
	fmt.Println("sesssion Email : " ,v )
	if  v == nil {
		flash := beego.NewFlash()
		flash.Error("Please login")
		flash.Store(&c.Controller)
		c.Redirect("/admin/login",302)
	}
	c.Data["Email"] = "243790881@qq.com"
	c.TplName = "admin/index.html"
}


func (c *AdminController) Login()  {

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
		qs := o.QueryTable("user")
		var userRes []*models.Users
		num,err := qs.Filter("email",Email).Filter("password",Passwored).All(&userRes)
		flask := beego.NewFlash()
		if err != nil || num < 1 {
			flask.Error("Email is not exsit!")
			flask.Store(&c.Controller)
			c.Ctx.Redirect(302,"/admin/login")
  			c.StopRun()
		}
		flask.Error("Email is not exsit!")
		flask.Store(&c.Controller)
		c.SetSession("email",Email)
		c.Redirect("/admin/",302)
	}

}

func (c *AdminController)Logout() {
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


func (c *AdminController) Register()  {
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

// @router /forgot-password [get]
func (c *AdminController) ForgotPassword()  {
	c.TplName = "admin/forgotpassword.html"
}

// @router /error [get]
func (c *AdminController) Error404()  {
	c.TplName = "admin/404.html"
}


type ImgInfo struct {
	Link string		`json:"link"`
}



func (c *AdminController) UploadImg() {
	file , h , err :=  c.GetFile("file")
	if err != nil{
		imginfo := ImgInfo{"error"}
		c.Data["json"] = &imginfo
		c.ServeJSON()
	}
	dirName := time.Now().Format("2006-01-02")
	pathName := "static/images/" +  dirName
	_, err = os.Stat(pathName)
	if os.IsNotExist(err) {
		os.Mkdir(pathName,os.ModePerm)
	}
	path := fmt.Sprintf("%s/%s",pathName , h.Filename)
	file.Close()
	c.SaveToFile("file",path)
	imginfo := ImgInfo{fmt.Sprintf("%s/%s",c.Ctx.Request.Header["Origin"][0],path)}
	c.Data["json"] = &imginfo
	c.ServeJSON()
}

// @router /blank [get]
func (c *AdminController) Blank()  {
	c.Data["Ls"] = "10056"
	c.TplName = "admin/blank.html"
}

// @router /charts [get]
func (c *AdminController) Charts()  {
	c.TplName = "admin/charts.html"
}

// @router /tables [get]
func (c *AdminController) Tables()  {
	c.TplName = "admin/tables.html"
}

// @router /article [get]
func (c *AdminController) Article()  {
	c.TplName = "admin/article.html"
}



