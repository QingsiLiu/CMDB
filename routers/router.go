package routers

import (
	"github.com/astaxie/beego"
	"magego/course-33/cmdb/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.UserController{})
}
