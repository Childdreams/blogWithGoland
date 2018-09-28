package controllers

import (
	"github.com/astaxie/beego/orm"
	"myproject/models"
	"fmt"
	"time"
	"os"
	"strings"
	"myproject/utils"
	"myproject/controllers/middleware"
)

type ArticleController struct {
	middleware.BaseController
}


func (c *ArticleController) AddArticle()  {
	if c.Ctx.Request.Method == "GET"{
		o := orm.NewOrm()
		var blogT []*models.Blogtype
		o.QueryTable("blogtype").OrderBy("-serialnum").All(&blogT)
		var contentInfo models.BbBlogContent
		ContentID , err := o.Insert(&contentInfo)
		if err != nil || ContentID <= 0{
			return
		}
		var article models.Blog
		article.Blog_content_id = int(ContentID)
		article.Title = "新建文章"
		article.Created = time.Now()
		article.Blog_content_last_update = time.Now()
		articleId , err := o.Insert(&article)
		if err != nil || articleId <= 0 {

		}
		var ArticleList utils.BlogContent
		c.Data["articleList"] = ArticleList
		c.Data["select"] =  blogT
		c.Data["articleID"] = articleId
		c.TplName = "admin/addArticle.html"
	}else {
		file , h , err :=  c.GetFile("view")
		Title := c.GetString("title")
		Type := c.GetString("type")
		Id ,_ := c.GetInt("ArticleID" , 0)
		if Id == 0 {

		}
		path := ""
		if err == nil {
			dirName := time.Now().Format("2006-01-02")
			pathName := "static/images/" +  dirName
			_, err = os.Stat(pathName)
			if os.IsNotExist(err) {
				os.Mkdir(pathName,os.ModePerm)
			}
			fmt.Println(h.Filename)
			fileType := strings.Split(h.Filename,".")
			fmt.Println(fileType)
			fileName := fmt.Sprintf("%s.%s",utils.EnMd5(h.Filename + dirName),fileType[len(fileType) - 1])
			path := fmt.Sprintf("%s/%s",pathName , fileName)
			defer file.Close()
			c.SaveToFile("file",path)
		}
		var updataInfo orm.Params
		updataInfo = make(orm.Params)
		updataInfo["title"] = Title
		updataInfo["type"] = Type
		if path != "" {
			updataInfo["views"] = path
		}
		orm.NewOrm().QueryTable("blog").Filter("id",Id).Update(
			updataInfo,
		)
		c.Redirect("AddArticle", 302)
	}
}

func (c *ArticleController)UploadArticle()  {
	body := c.GetString("body")
	articleId := c.GetString("ArticleID")
	o := orm.NewOrm()
	var blog models.Blog
	o.QueryTable("blog").Filter("id",articleId).One(&blog)
	num , err := o.QueryTable("bb_blog_content").Filter("id",blog.Blog_content_id).Update(
		orm.Params{
			"content" : body,
		},
	)
	o.QueryTable("blog").Filter("id",articleId).Update(
		orm.Params{
			"blog_content_last_update":time.Now(),
		},
	)
	if err != nil {
		fmt.Println(err)
		fmt.Println(num)
	}
	jsonInfo := make(map[string]string)
	jsonInfo["errorCode"] = "200"
	c.Data["json"] = &jsonInfo
	c.ServeJSON()
}

func (c *ArticleController)ArticleList(){
	var ArticleList []utils.BlogContent
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("t1.*","t2.content").
		From("blog as t1").
		LeftJoin("bb_blog_content as t2").
		On("t1.blog_content_id = t2.id")
	sql := qb.String()
	orm.NewOrm().Raw(sql).QueryRows(&ArticleList)
	c.Data["articleList"] = &ArticleList
	c.TplName = "admin/ArtileList.html"
}

func (c *ArticleController)ModArticle (){
	id := c.GetString("id")
	var blogT []*models.Blogtype
	o := orm.NewOrm()
	o.QueryTable("blogtype").OrderBy("-serialnum").All(&blogT)
	var ArticleList utils.BlogContent
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("t1.*","t2.content").
		From("blog as t1").
		LeftJoin("bb_blog_content as t2").
		On("t1.blog_content_id = t2.id").
		Where("t1.id = ?")
	sql := qb.String()
	orm.NewOrm().Raw(sql,id).QueryRow(&ArticleList)
	fmt.Println(ArticleList)
	c.Data["select"] =  blogT
	c.Data["articleID"] = id
	c.Data["articleList"] = ArticleList
	c.TplName = "admin/addArticle.html"
}

func (c *ArticleController)DelArticle (){
	id := c.GetString("id")
	var delArticle models.Blog
	o := orm.NewOrm()
	o.QueryTable("blog").Filter("id",id).One(&delArticle)
	o.QueryTable("bb_blog_content").Filter("id",delArticle.Blog_content_id).Delete()
	o.QueryTable("blog").Filter("id",id).Delete()
	c.Redirect("ArticleList", 302)
}