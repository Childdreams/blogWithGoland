package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	_ "myproject/models"
	"myproject/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"myproject/utils"
	"strconv"
)

type RbacController struct {
	BaseController
}

//Permissions Controller
func (c *RbacController) AddPermissions()  {
	if c.Ctx.Request.Method == "GET"{
		o := orm.NewOrm()
		qs := o.QueryTable("permissions")
		var permissions []*models.Permissions
		_,err := qs.All(&permissions)
		if err != nil{

		}
		permissionTree := utils.GetAllPerm(permissions,0,0)
		fmt.Println(permissionTree)
		c.Data["tree"] = permissionTree
		c.TplName = "admin/AddPermissions.html"
	}else {
		flask := beego.NewFlash()
		Name := c.GetString("name")
		o :=  orm.NewOrm()
		display_name := c.GetString("display_name")
		description := c.GetString("description")
		is_menu :=c.GetString("is_menu")
		Parent_id := c.GetString("parent_id")
		permission :=  new(models.Permissions)
		fmt.Println(display_name,description,Name)
		if Name == "" || display_name == "" {
			flask.Error("Name or display_name not empty")
			c.Redirect("addpermissions",302)
		}
		permission.Name = Name
		permission.Display_name = display_name
		permission.Description = description
		is_menu_int ,_ := strconv.Atoi(is_menu)
		permission.Is_menu = is_menu_int
		parent_id_int ,_ :=strconv.Atoi(Parent_id)
		permission.Parent_id = parent_id_int
		permission.Created_at = time.Now()
		permission.Updated_at = time.Now()
		_,err := o.Insert(permission)
		if err != nil{
			flask.Error("Insert error")
		}
		flask.Notice("Creating a successful")
		flask.Store(&c.Controller)
		c.Redirect("addpermissions",302)
	}
}

func (c *RbacController) PermissionsLists (){
	o := orm.NewOrm()
	qs := o.QueryTable("permissions")
	var permissions []*models.Permissions
	_,err := qs.All(&permissions)
	if err != nil {

	}
	permissionTree := utils.GetAllPerm(permissions,0,0)
	c.Data["list"] = permissionTree
 	c.TplName = "admin/permissionList.html"
}

func (c *RbacController) DelPermission()  {
	flash := beego.NewFlash()
	Id := c.GetString("id","0")
	if Id == "0"{
		flash.Error("Delete error!!!")
		c.Redirect("permissionsLists" , 302)
	}
	o := orm.NewOrm()
	cond := orm.NewCondition()
	condPermission := cond.And("id",Id).Or("parent_id",Id)
	qs := o.QueryTable("permissions")
	cnt ,_ := qs.SetCond(condPermission).Count()
	if cnt != 1 {
		flash.Error("Delete error!!!")
		c.Redirect("permissionsLists" , 302)
	}
	fmt.Println("this is delete :",cnt)
	num , _ := qs.SetCond(condPermission).Delete()
	if num <= 0 {
		flash.Error("Delete error!!!")
		c.Redirect("permissionsLists" , 302)
	}
	flash.Notice("DELETE SUCCESSFULL")
	flash.Store(&c.Controller)
	c.Redirect("permissionsLists" , 302)
}

func (c *RbacController) ModPermission () {
	if c.Ctx.Request.Method == "GET" {
		Id := c.GetString("id")
		o := orm.NewOrm()
		cond := orm.NewCondition()
		condPermission := cond.And("id",Id)
		var permissionsInfo models.Permissions

		qs := o.QueryTable("permissions")
		qs.SetCond(condPermission).One(&permissionsInfo)

		qs = o.QueryTable("permissions")
		var permissions []*models.Permissions
		_,err := qs.All(&permissions)
		if err != nil {

		}
		permissionTree := utils.GetAllPerm(permissions,0,0)
		c.Data["tree"] = permissionTree
		c.Data["permissionsInfo"] = permissionsInfo
		c.TplName = "admin/ModPermissions.html"
	}else {
		flask := beego.NewFlash()
		Name := c.GetString("name")
		Id := c.GetString("id")
		o :=  orm.NewOrm()
		display_name := c.GetString("display_name")
		description := c.GetString("description")
		Parent_id := c.GetString("parent_id")
		is_menu :=c.GetString("is_menu")
		parent_id_int ,_ :=strconv.Atoi(Parent_id)
		fmt.Println(Name,Id,display_name,description,Parent_id,is_menu,parent_id_int)
		if Name == "" || display_name == "" || Id == ""{
			flask.Error("Name or display_name not empty")
			c.Redirect("ModPermission",302)
		}
		num,_ := o.QueryTable("permissions").Filter("id",Id).Update(
			orm.Params{
				"name":Name,"display_name":display_name,"description":description,"is_menu":is_menu,"parent_id":parent_id_int,"updated_at":time.Now(),
			})
		if num < 1 {
			flask.Error("Nothing modified")
		}else {
			flask.Notice("Modified success")
		}
		flask.Store(&c.Controller)
		c.Redirect("permissionsLists" , 302)
	}
}





//Role Controller

func (c *RbacController)RoleList()  {
	// use Structure query
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("roles.*","GROUP_CONCAT(permissions.display_name) as dname").
		From("roles").
		LeftJoin("permission_role").
		On("permission_role.role_id = roles.id").
		LeftJoin("permissions").
		On("permissions.id = permission_role.permission_id").
		GroupBy("roles.id")
	sql := qb.String()
	o := orm.NewOrm()
	var role []utils.RolesSQ
	o.Raw(sql).QueryRows(&role)
	c.Data["roles"] = role
	c.TplName = "admin/Role/RoleList.html"

}

func (c *RbacController)AddRole()  {

	if c.Ctx.Request.Method == "GET"{
		tree := utils.GetAllPerInfo()
		c.Data["tree"] = tree
		c.TplName = "admin/Role/AddRole.html"
	}else {
		RequestData := c.Input()
		RoleName := RequestData["name"][0]
		display_name := RequestData["display_name"][0]
		description := RequestData["description"][0]
		permissionID := RequestData["permissionID"]
		o := orm.NewOrm()
		var role  models.Roles
		role.Name = RoleName
		role.Display_name = display_name
		role.Description = description
		role.Created_at = time.Now()
		role.Updated_at = time.Now()
		id ,err:= o.Insert(&role)
		flash := beego.NewFlash()

		if id < 1 ||err != nil {
			flash.Error("add new Roles Erroe")
			c.Redirect("addrole",302)
		}

		permission_role := []models.PermissionRole{}

		for _,m := range permissionID{
			m_int ,_ :=strconv.Atoi(m)
			insertype := models.PermissionRole{PermissionId:m_int,RoleId:id}
			permission_role = append(permission_role,insertype)
		}

		successNums, _ := o.InsertMulti(100, permission_role)

		if successNums < 1{
			flash.Notice("add new Roles Successfull")
		}

		flash.Store(&c.Controller)
		c.Redirect("addrole",302)
	}
}

func (c *RbacController)ModRole()  {
	if c.Ctx.Request.Method == "GET" {
		Id := c.GetString("id")
		cond :=orm.NewCondition()
		condRoles := cond.And("id",Id)
		var resRoles models.Roles
		err := orm.NewOrm().QueryTable("roles").SetCond(condRoles).One(&resRoles)
		if err == orm.ErrNoRows {
			// 没有找到记录
			fmt.Printf("Not row found roles")
		}
		condPerRole := cond.And("role_id",Id)
		var resPerRole  []*models.PermissionRole
		num ,_ := orm.NewOrm().QueryTable("permissionRole").SetCond(condPerRole).All(&resPerRole)
		if num < 1 {
			fmt.Printf("Not row found permissionRole")
		}
		c.Data["roles"] = resRoles
		var perRoles []int
		for _ ,v := range resPerRole{
			perRoles = append(perRoles, v.PermissionId)
		}
		c.Data["preRoles"] = perRoles

		tree := utils.GetAllPerInfo()
		c.Data["tree"] = tree
		c.TplName = "admin/Role/ModRole.html"
	}else {
		RequestData := c.Input()
		Id := RequestData["id"][0]
		RoleName := RequestData["name"][0]
		display_name := RequestData["display_name"][0]
		description := RequestData["description"][0]
		permissionID := RequestData["permissionID"]
		o := orm.NewOrm()
		num,_ := o.QueryTable("roles").Filter("id",Id).Update(
			orm.Params{
				"name":RoleName,"display_name":display_name,"description":description,"updated_at":time.Now(),
			})
		flash := beego.NewFlash()
		if num < 1  {
			flash.Error("add new Roles Erroe")
			c.Redirect("addrole",302)
		}
		o.QueryTable("permissionRole").Filter("role_id",Id).Delete()
		permission_role := []models.PermissionRole{}
		id , _:= strconv.ParseInt(Id,10,64)
		for _,m := range permissionID{
			m_int ,_ :=strconv.Atoi(m)
			insertype := models.PermissionRole{PermissionId:m_int,RoleId:id}
			permission_role = append(permission_role,insertype)
		}

		successNums, _ := o.InsertMulti(100, permission_role)

		if successNums < 1{
			flash.Notice("add new Roles Successfull")
		}

		flash.Store(&c.Controller)
		c.Redirect("rolelist",302)
	}
}

func (c *RbacController)DelRole()  {
	flash := beego.NewFlash()
	Id := c.GetString("id","0")
	if Id == "0"{
		flash.Error("Delete error!!!")
		c.Redirect("rolelist" , 302)
	}
	o := orm.NewOrm()
	cond := orm.NewCondition()
	condRoles := cond.And("id",Id)
	qs := o.QueryTable("roles")

	num , _ := qs.SetCond(condRoles).Delete()
	o.QueryTable("permissionRole").Filter("role_id",Id).Delete()
	if num <= 0 {
		flash.Error("Delete error!!!")
		c.Redirect("rolelist" , 302)
	}
	flash.Notice("DELETE SUCCESSFULL")
	flash.Store(&c.Controller)
	c.Redirect("rolelist" , 302)
}

// User Info
func (c *RbacController)UserList()  {

}

func (c *RbacController)AddUser()  {
	c.Ctx.WriteString("ceshi")
	c.StopRun()
}
func (c *RbacController)ModUser()  {

}
func (c *RbacController)DelUser()  {

}