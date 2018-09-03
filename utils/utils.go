package utils

import (
	"crypto/md5"
	"encoding/hex"
	"myproject/models"
	"github.com/astaxie/beego/orm"
)

func EnMd5(EncryptedStr  string) string  {
	h := md5.New()
	h.Write([]byte(EncryptedStr))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

type PermissionsCount struct {
	*models.Permissions
	Count int
}



func GetAllPerm(permissions []*models.Permissions , id int  ,count int) []PermissionsCount{
	res  := []PermissionsCount{}
	res = GetTree(res,permissions,id,count)
	return res
}

func GetTree(res []PermissionsCount ,permissions []*models.Permissions , id int  ,count int ) []PermissionsCount{
	for _ ,permission := range permissions{
		if permission.Parent_id == id {
			per := PermissionsCount{permission,count}
			res = append(res, per)
			res = GetTree(res, permissions,permission.Id,count+1)
		}
	}
	return res
}

func GetAllPerInfo() []PermissionsCount {
	res  := []PermissionsCount{}
	o := orm.NewOrm()
	var permissionsTree []*models.Permissions
	qs := o.QueryTable("permissions")
	_,err := qs.All(&permissionsTree)
	if err != nil {

	}
	res = GetTree(res,permissionsTree,0,0)
	return res
}