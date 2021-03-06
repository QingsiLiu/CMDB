package controllers

import (
	"github.com/astaxie/beego"
	"magego/course-33/cmdb/base/controllers/auth"
	"magego/course-33/cmdb/forms"
	"magego/course-33/cmdb/services"
	"net/http"
)

// UserController 用户管理控制器
type UserController struct {
	auth.LayoutController
}

// Query 查询用户
func (u *UserController) Query() {
	beego.ReadFromRequest(&u.Controller)

	q := u.GetString("q")

	u.Data["users"] = services.UserService.Query(q)
	u.Data["q"] = q
	u.TplName = "user/query.html"
	u.Data["title"] = "用户列表"
}

// New 新建用户
func (u *UserController) New() {
	form := &forms.UserModifyForm{}

	if u.Ctx.Input.IsPost() {
		if err := u.ParseForm(form); err == nil {
			services.UserService.New(form)
			u.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
		}
	}
	//u.Data["form"] = form
	u.Data["xsrf_token"] = u.XSRFToken()
	u.TplName = "user/new.html"
	u.Data["title"] = "用户新建"
	u.LayoutSections["SectionStyle"] = "user/new_style.html"
}

// Modify 修改用户
func (u *UserController) Modify() {
	// 假设当前用户不能修改其他人的信息
	/*if 1==1{
		u.Abort("NotPermission")
		return
	}*/

	form := &forms.UserModifyForm{}

	// Get获取数据，Post修改用户提交数据
	if u.Ctx.Input.IsPost() {
		if err := u.ParseForm(form); err == nil {
			// 验证数据
			services.UserService.Modify(form)

			// 存储消息
			flash := beego.NewFlash()
			flash.Set("notice", "修改用户信息成功")
			flash.Store(&u.Controller)

			u.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
		}

	} else if pk, err := u.GetInt("pk"); err == nil {
		if user := services.UserService.GetByID(pk); user != nil {
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
	u.Data["xsrf_token"] = u.XSRFToken()
	u.TplName = "user/modify.html"
	u.Data["title"] = "用户修改"
}

// Delete 删除用户
func (u *UserController) Delete() {
	if pk, err := u.GetInt("pk"); err == nil {
		services.UserService.Delete(pk)
	}

	u.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
}
