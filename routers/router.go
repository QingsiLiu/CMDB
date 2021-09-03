package routers

import (
	"github.com/astaxie/beego"
	"magego/course-33/cmdb/controllers"
)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.PasswordController{})

	beego.Router("/", &controllers.HomeController{}, "*:Index")

	beego.ErrorController(&controllers.ErrorController{})
}
