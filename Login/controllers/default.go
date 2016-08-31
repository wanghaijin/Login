package controllers

import (
	"Login/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) IndexGet() {
	c.TplName = "index.tpl"
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
}

func (c *MainController) IndexPost() {
	username := c.GetString("username")
	password := c.GetString("password")
	email := c.GetString("email")

	user := models.User{}
	user.Email = email
	user.Username = username
	user.Password = password

	if models.CheckUser(&user) {
		c.Data["json"] = "用户已存在!"
		c.ServeJSON()
	} else {
		if username != nil; password != nil {
			_, err := models.AddUser(&user)
			if err != nil {
				c.Data["json"] = "添加失败"
				c.ServeJSON()
			} else {
				c.Data["json"] = "添加成功"
				c.ServeJSON()
			}
		} else {
			c.Data["json"] = "用户名或密码为空，无法添加!"
			c.ServeJSON()
		}

	}

}

func (c *MainController) LoginGet() {
	c.TplName = "login.tpl"
}

func (c *MainController) LoginPost() {
	username := c.GetString("username")
	password := c.GetString("password")
	user, err := models.GetUserByUsername(username)
	if err != nil {
		c.Ctx.WriteString("用户名错误!")
	}
	if user.Password == password {
		c.Ctx.WriteString("登录成功!")
	} else {
		c.Ctx.WriteString("密码错误!")
	}
}

func (c *MainController) RegeditGet() {
	c.TplName = "regedit.tpl"
}

func (c *MainController) RegeditPost() {

}
