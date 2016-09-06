package controllers

import (
	"Login/models"
	//"fmt"
	"github.com/astaxie/beego"

)

type MainController struct {
	beego.Controller
}

func (c *MainController) IndexGet() {
	c.TplName = "index.html"
}

func (c *MainController) IndexPost() {

	username := c.GetString("username")
	password := c.GetString("password")
	email := c.GetString("email")


	user, errU := models.GetUserByUsername(username)
	if errU != nil {
		c.Data["json"] = "用户名错误!"
		c.ServeJSON()
	}
	_,errE:=models.GetUserByEmail(email)
	if  errE!=nil{
		c.Data["json"]="邮箱错误!"
		c.ServeJSON()

	}
	if user.Password == password {
		c.Data["json"] = "OK"
		c.ServeJSON()
	} else {
		c.Data["json"] = "密码错误!"
		c.ServeJSON()
	}

}



func (c *MainController) RegeditGet() {

	c.TplName = "regedit.html"
}

func (c *MainController) RegeditPost() {

	username := c.GetString("username")
	password := c.GetString("password")
	email := c.GetString("email")

	user := models.User{}
	user.Email = email
	user.Username = username
	user.Password = password

	if models.CheckUserName(&user) {
		c.Data["json"] = "用户名已存在!"
		c.ServeJSON()

	} else if models.CheckUserEmail(&user){
		c.Data["json"]="邮箱已注册"
		c.ServeJSON()

	}else{


		_, err := models.AddUser(&user)
		if err != nil {
			c.Data["json"]="添加失败"
			c.ServeJSON()


		} else {
			c.Data["json"]="添加成功,请登录！"
			c.ServeJSON()
		}

	}

}
