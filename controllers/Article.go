package controllers

import (
	"github.com/astaxie/beego/orm"
	"myproject/models"
	"fmt"
	"time"
)

type ArticleController struct {
	BaseController
}


func (c *ArticleController) AddArticle()  {
	if c.Ctx.Request.Method == "GET"{
		o := orm.NewOrm()
		var blogT []*models.Blogtype
		o.QueryTable("blogtype").OrderBy("-serialnum").All(&blogT)
		fmt.Println(blogT)
		var contentInfo models.BbBlogContent
		ContentID , err := o.Insert(&contentInfo)
		if err != nil || ContentID <= 0{
			return
		}
		var article models.Blog
		article.Blog_content_id = int(ContentID)
		article.Title = "新建文章"
		article.Created = time.Now()
		articleId , err := o.Insert(&article)
		if err != nil || articleId <= 0 {

		}
		c.Data["select"] =  blogT
		c.Data["articleID"] = articleId
		c.TplName = "admin/addArticle.html"

	}else {
		fmt.Println(c.Input())
	}
}

func (c *ArticleController)UploadArticle()  {
	body := c.GetString("body")
	articleId := c.GetString("ArticleID")
	o := orm.NewOrm()
	o.QueryTable("blog")
}