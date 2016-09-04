package controllers

import (
	"Login/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) IndexGet() {
	c.TplName = "index.html"
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.html"
}

func (c *MainController) IndexPost() {
	username := c.GetString("username")
	password := c.GetString("password")
	user, err := models.GetUserByUsername(username)
	if err != nil {
		//c.Ctx.Redirect(302,"/")
		//c.TplName="index.html"
		c.Data["json"] = "用户名错误!"
		//c.Ctx.Output
		c.ServeJSON()
	}
	if user.Password == password {
		//c.Ctx.Redirect(302,"login.html")
		c.Data["json"] = "OK"
		c.ServeJSON()
		//c.TplName="login.html"
	} else {
		c.Ctx.Redirect(302,"/")

		//c.TplName="index.html"
		c.Data["json"] = "密码错误!"
		c.ServeJSON()
	}

}

func (c *MainController) LoginGet() {
	c.TplName = "login.html"
}

func (c *MainController) LoginPost() {

}

func (c *MainController) RegeditGet() {
	//c.Data["error"]=c.Data["regedit"]
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
		//c.ServeJSON()
		c.Ctx.Redirect(302,"regedit.html")
	} else if models.CheckUserEmail(&user){
		c.Data["json"]="邮箱已注册"
		c.Ctx.Redirect(302,"regedit.html")
	}else{

		// if username != nil; password != nil {
		_, err := models.AddUser(&user)
		if err != nil {
			c.Data["json"] = "添加失败"
			//c.ServeJSON()
			c.Ctx.Redirect(302,"regedit.html")
			//c.TplName="index.html"

		} else {
			c.Data["json"] = "添加成功"
			//c.ServeJSON()
			//c.Ctx.Output.JSON(c.Data["json"],true)
			c.Ctx.Redirect(302,"/")

			//c.TplName="login.html"

		}
		// } else {
		// 	c.Data["json"] = "用户名或密码为空，无法添加!"

		// }

	}
	//c.ServeJSON()

}
