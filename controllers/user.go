package controllers

import (
	"github.com/astaxie/beego"
	"magego/course-33/cmdb/base/controllers/base"
	"magego/course-33/cmdb/models"
	"net/http"
)

// 用户管理控制器
type UserController struct {
	base.BaseController
}

// Query 查询用户
func (u *UserController) Query() {
	sessionUser := u.GetSession("user")
	if sessionUser == nil {
		u.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
		return
	}

	q := u.GetString("q")

	users := models.QueryUser(q)
	u.Data["users"] = users
	u.Data["q"] = q
	u.TplName = "user/query.html"
}
