package controllers

import (
	"github.com/astaxie/beego"
	"magego/course-33/cmdb/base/controllers/auth"
	"magego/course-33/cmdb/forms"
	"magego/course-33/cmdb/models"
	"net/http"
)

// UserController 用户管理控制器
type UserController struct {
	auth.AuthorizationController
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

// New 新建用户
func (u *UserController) New() {
	form := &forms.UserModifyForm{}

	if u.Ctx.Input.IsPost() {
		if err := u.ParseForm(form); err == nil {
			models.NewUser(form)
			u.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
		}
	}
	//u.Data["form"] = form
	u.TplName = "user/new.html"
}

// Modify 修改用户
func (u *UserController) Modify() {
	form := &forms.UserModifyForm{}

	// Get获取数据，Post修改用户提交数据
	if u.Ctx.Input.IsPost() {
		if err := u.ParseForm(form); err == nil {
			// 验证数据
			models.ModifyUser(form)
			u.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
		}

	} else if pk, err := u.GetInt("pk"); err == nil {
		if user := models.GetUserByID(pk); user != nil {
			form.ID = user.ID
			form.StaffId = user.StaffID
			form.NickName = user.Nickname
			form.Gender = user.GenderText()
			form.Department = user.Department
			form.Email = user.Email
			form.Tel = user.Tel
			form.Name = user.Name
			form.Addr = user.Addr
			form.Status = user.StatusText()
		}
	}

	u.Data["form"] = form
	u.TplName = "user/modify.html"
}

// Delete 删除用户
func (u *UserController) Delete() {
	form := &forms.UserModifyForm{}
	if err := u.ParseForm(form); err == nil {
		if pk, err := u.GetInt("pk"); err == nil && u.LoginUser.ID != pk {
			models.DeleteUser(form)
		}
	}

	u.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
}
