package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int    `orm:"cloumn(id);auto" `
	Username string `orm:"cloumn(username);size(20)"`
	Password string `orm:"cloumn(password);size(20)"`
	Email    string `orm:"cloumn(email);size(20)"`
}

func (c *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

//添加用户
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()

	if _, err = o.Insert(m); err != nil {
		return
	}
	return
}

//检查用户名是否存在
func CheckUserName(m *User) bool {
	o := orm.NewOrm()
	return o.QueryTable(new(User)).Filter("username", m.Username).Exist()
}

//检查email是否存在
func CheckUserEmail(m *User) bool {
	o := orm.NewOrm()
	return o.QueryTable(new(User)).Filter("email", m.Email).Exist()
}

//根据用户名获取用户信息
func GetUserByUsername(username string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Username: username}
	if err = o.Read(v, "Username"); err != nil {
		return
	}
	return
}

//根据email获取用户信息
// func GetUserByEmail(email string) (v *User, err error) {
// 	o := orm.NewOrm()
// 	v = &User{Email: email}
// 	if err = o.Read(v, "Email"); err != nil {
// 		return
// 	}
// 	return
// }
