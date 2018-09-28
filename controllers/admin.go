package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	_ "myproject/models"
	"time"
	"fmt"
	"github.com/astaxie/beego"
	"myproject/utils"
	"os"
	"strings"
	"myproject/controllers/middleware"
)

type AdminController struct {
	middleware.BaseController
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
	v := c.GetSession("UserInfo")
	if  v == nil {
		flash := beego.NewFlash()
		flash.Error("Please login")
		flash.Store(&c.Controller)
		c.Redirect("/admin/login",302)
	}
	c.Data["Email"] = "243790881@qq.com"
	c.TplName = "admin/index.html"
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
	fileType := strings.Split(h.Filename,".")
	fileName := fmt.Sprintf("%s.%s",utils.EnMd5(h.Filename + dirName),fileType[len(fileType) - 1])
	path := fmt.Sprintf("%s/%s",pathName , fileName)
	defer file.Close()
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



