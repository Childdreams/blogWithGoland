package utils

import (
	"crypto/md5"
	"encoding/hex"
	"myproject/models"
	"github.com/astaxie/beego/orm"
	"sort"
	"fmt"
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


func ForAdd(num int ,in string)(out string){
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