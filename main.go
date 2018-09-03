package main

import (
	_ "myproject/models"
	_ "myproject/routers"
	"github.com/astaxie/beego"
	"myproject/controllers"
	"github.com/astaxie/beego/orm"
	"sort"
	"fmt"
)

func forAdd(num int ,in string)(out string){
	str := ""
	for i:= 1 ; i<=num;i++{
		str += in
	}
	return str
}

func In_array(arrInt []int , target int)(bool){
	sort.Ints(arrInt)
	i := sort.Search(len(arrInt), func(i int) bool {
		return arrInt[i] >= target
	})
	fmt.Println("this is test")
	if i<len(arrInt) && arrInt[i] == target { //这里可以采用 strings.EqualFold(arrString[i],target)
		return true
	}else {
		return  false
	}
}
func main() {
	beego.ErrorController(&controllers.AdminController{})
	if beego.AppConfig.String("debug") == "true" {
		orm.Debug = true
		beego.BConfig.EnableErrorsShow = true
		beego.BConfig.Log.AccessLogs = true
	}
	beego.AddFuncMap("hi",forAdd)
	beego.AddFuncMap("In_array",In_array)

	beego.AddTemplateExt("html")
	beego.Run()
}

