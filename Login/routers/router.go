package routers

import (
	"Login/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:IndexGet")
	beego.Router("/", &controllers.MainController{}, "post:IndexPost")
	beego.Router("/login", &controllers.MainController{}, "get:LoginGet")
	beego.Router("/login", &controllers.MainController{}, "post:LoginPost")
	beego.Router("/regedit", &controllers.MainController{}, "get:RegeditGet")
	beego.Router("/regedit", &controllers.MainController{}, "post:RegeditPost")
}
