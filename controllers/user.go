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
	q := u.GetString("q")

	users := models.QueryUser(q)
	u.Data["users"] = users
	u.Data["q"] = q
	u.TplName = "user/query.html"
}
