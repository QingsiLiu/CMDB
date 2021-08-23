package controllers

import (
	"github.com/astaxie/beego"
	"magego/course-33/cmdb/models"
)

// 用户管理控制器
type UserController struct {
	beego.Controller
}

// Query 查询用户
func (u *UserController) Query() {
	users := models.QueryUser("")
	u.Data["users"] = users
	u.TplName = "user/query.html"
}
