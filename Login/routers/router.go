package routers

import (
	"Login/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:IndexGet")
	beego.Router("/", &controllers.MainController{}, "post:IndexPost")
	beego.Router("/register", &controllers.MainController{}, "get:RegisterGet")
	beego.Router("/register", &controllers.MainController{}, "post:RegisterPost")
	beego.Router("/cache", &controllers.MainController{}, "get:CacheGet")
}
