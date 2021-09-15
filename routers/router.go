package routers

import (
	"github.com/astaxie/beego"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"magego/course-33/cmdb/controllers"
	v1 "magego/course-33/cmdb/controllers/api/v1"
	"magego/course-33/cmdb/filters"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeExec, filters.BeforeExecute)
	beego.InsertFilter("/*", beego.AfterExec, filters.AfterExecute, false)

	beego.Handler("/metrics", promhttp.Handler())

	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.PasswordController{})
	beego.AutoRouter(&controllers.TaskController{})

	// prometheus
	beego.AutoRouter(&controllers.NodeController{})
	beego.AutoRouter(&controllers.JobController{})
	beego.AutoRouter(&controllers.TargetController{})

	// /v1/
	v1Prom := beego.NewNamespace("/v1", beego.NSAutoRouter(&v1.PrometheusController{}))
	beego.AddNamespace(v1Prom)

	beego.Router("/", &controllers.HomeController{}, "*:Index")

	beego.ErrorController(&controllers.ErrorController{})
}
