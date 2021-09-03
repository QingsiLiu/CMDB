package controllers

import "magego/course-33/cmdb/base/controllers/base"

// ErrorController 错误处理控制器
type ErrorController struct {
	base.BaseController
}

//Error 开头的方法

// Error404 处理404错误
func (e *ErrorController) Error404() {
	e.TplName = "error/404.html"
}

func (e *ErrorController) ErrorNotPermission() {
	e.TplName = "error/not_permission.html.html"
}
