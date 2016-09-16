package models

import (
	"github.com/astaxie/beego/orm"
)

func init(){
	orm.RegisterDataBase("default","mysql","root:6917@/mydb?charset=utf8&loc=Asia%2FShanghai")
}
