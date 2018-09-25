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
		qs := o.QueryTable("permissions").OrderBy("-serialnum")
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
		Serialnum := c.GetString("serialnum")
		permission :=  new(models.Permissions)
		fmt.Println(display_name,description,Name)
		if Name == "" || display_name == "" {
			flask.Error("Name or display_name not empty")
			c.Redirect("addpermissions",302)
		}
		permission.Name = Name
		permission.Display_name = display_name
		permission.Description = description
		Serialnum_int ,_ := strconv.Atoi(Serialnum)
		permission.Serialnum = Serialnum_int
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
		serialnum :=c.GetString("serialnum")
		serialnum_int ,_ :=strconv.Atoi(serialnum)
		parent_id_int ,_ :=strconv.Atoi(Parent_id)
		fmt.Println(Name,Id,display_name,description,Parent_id,is_menu,parent_id_int)
		if Name == "" || display_name == "" || Id == ""{
			flask.Error("Name or display_name not empty")
			c.Redirect("ModPermission",302)
		}
		num,_ := o.QueryTable("permissions").Filter("id",Id).Update(
			orm.Params{
				"name":Name,"display_name":display_name,"description":description,"is_menu":is_menu,"parent_id":parent_id_int,"updated_at":time.Now(),"serialnum":serialnum_int,
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
	DB ,_ := orm.NewQueryBuilder("mysql")
	DB.Select("users.id","users.name","users.email","users.remember_token","users.created_at","users.updated_at","GROUP_CONCAT(roles.display_name) as dname").
		From("users").
		LeftJoin("role_user").
		On("role_user.user_id = users.id").
		LeftJoin("roles").
		On("roles.id = role_user.role_id").
		GroupBy("users.id")
	var userdata []*utils.UserSQL
	orm.NewOrm().Raw(DB.String()).QueryRows(&userdata)
	c.Data["tree"] = userdata
	c.TplName = "admin/User/UserList.html"
}

func (c *RbacController)AddUser()  {
	if c.Ctx.Request.Method == "GET"{
		var rolesInfo []*models.Roles
		orm.NewOrm().QueryTable("roles").All(&rolesInfo)
		c.Data["tree"] = rolesInfo
		c.TplName = "admin/User/AddUser.html"
	}else {
		flash := beego.NewFlash()
		RequestData := c.Input()
		Name := RequestData["name"][0]
		Email := RequestData["email"][0]
		Password := RequestData["password"][0]
		ConfirmPassword := RequestData["confirmPassword"][0]
		RolesID := RequestData["rolesID"]
		var userinfo models.Users
		userinfo.Name = Name
		userinfo.Email = Email
		userinfo.CreatedAt = time.Now()
		userinfo.UpdatedAt = time.Now()
		fmt.Println(Password,ConfirmPassword)
		if Password == ConfirmPassword{
			flash.Error("Password inconsistency")
			flash.Store(&c.Controller)
			c.Redirect("AddUser",302)
		}
		userinfo.Password = utils.EnMd5(Password)
		id , _ := orm.NewOrm().Insert(&userinfo)
		fmt.Println(userinfo)
		if id < 1 {
			flash.Error("Created Fail")
			flash.Store(&c.Controller)
			c.Redirect("AddUser",302)
		}

		var UserRoles []models.RoleUser
		for _,m := range RolesID{
			m_int ,_ :=strconv.Atoi(m)
			insertype := models.RoleUser{RoleId:m_int,UserId:int(id)}
			UserRoles = append(UserRoles,insertype)
		}
		successNums, _ := orm.NewOrm().InsertMulti(100,UserRoles)
		if successNums < 1{
			flash.Notice("add new Roles Successfull")
		}
		flash.Store(&c.Controller)
		c.Redirect("AddUser",302)
	}
}
func (c *RbacController)ModUser()  {
	if c.Ctx.Request.Method == "POST"{
		flash := beego.NewFlash()
		RequestData := c.Input()
		Id := RequestData["id"][0]
		Name := RequestData["name"][0]
		Email := RequestData["email"][0]
		Password := RequestData["password"][0]
		ConfirmPassword := RequestData["confirmPassword"][0]
		RolesID := RequestData["rolesID"]
		fmt.Println(Password,ConfirmPassword)
		if Password != "" && Password == ConfirmPassword{
			flash.Error("Password inconsistency")
			flash.Store(&c.Controller)
			c.Redirect("userList",302)
		}
		params := orm.Params{"name":Name,"email":Email,"updated_at":time.Now()}
		if Password != ""{
			params = orm.Params{"name":Name,"password":utils.EnMd5(Password),"email":Email,"updated_at":time.Now()}
		}
		o := orm.NewOrm()
		o.QueryTable("users").Filter("id",Id).Update(params)
		o.QueryTable("role_user").Filter("user_id",Id).Delete()

		var UserRoles []models.RoleUser
		id , _:= strconv.ParseInt(Id,10,64)
		for _,m := range RolesID{
			m_int ,_ :=strconv.Atoi(m)
			insertype := models.RoleUser{RoleId:m_int,UserId:int(id)}
			UserRoles = append(UserRoles,insertype)
		}
		successNums, _ := orm.NewOrm().InsertMulti(100,UserRoles)
		if successNums < 1{
			flash.Notice("add new Roles Successfull")
		}
		flash.Store(&c.Controller)
		c.Redirect("userList",302)
	}else {
		Id := c.GetString("id" , "0")
		flash := beego.NewFlash()
		if Id == "0" {
			flash.Error("Server Error!")
			flash.Store(&c.Controller)
			c.Redirect("userList", 302)
		}
		var user models.Users
		orm.NewOrm().QueryTable("users").Filter("id",Id).One(&user)
		var roles []models.RoleUser
		orm.NewOrm().QueryTable("role_user").Filter("user_id",Id).All(&roles)
		var roleInt []int
		for _,v := range roles{
			roleInt = append(roleInt, v.RoleId)
		}
		var rolesInfo []*models.Roles
		orm.NewOrm().QueryTable("roles").All(&rolesInfo)
		c.Data["tree"] = rolesInfo
		c.Data["userInfo"] = user
		c.Data["roleInfo"] = roleInt
		c.TplName = "admin/User/ModUser.html"
	}

}
func (c *RbacController)DelUser()  {
	Id := c.GetString("id" , "0")
	flash := beego.NewFlash()
	if Id == "0" {
		flash.Error("Server Error!")
		flash.Store(&c.Controller)
		c.Redirect("userList", 302)
	}
	orm.NewOrm().QueryTable("users").Filter("id",Id).Delete()
	orm.NewOrm().QueryTable("role_user").Filter("user_id",Id).Delete()
	flash.Notice("SUCCESS")
	flash.Store(&c.Controller)
	c.Redirect("userList",302)
}