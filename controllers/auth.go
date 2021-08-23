package controllers

import (
	"github.com/astaxie/beego"
	"magego/course-33/cmdb/base"
	"magego/course-33/cmdb/forms"
	"magego/course-33/cmdb/models"
	"net/http"
)

//负责认证的认证控制器
type AuthController struct {
	beego.Controller
}

//认证登录
func (a *AuthController) Login() {
	form := &forms.LoginForm{}
	errs := base.NewErrors()
	//Get请求直接加载页面
	//Post请求记性数据验证（成功/失败）
	if a.Ctx.Input.IsPost() {
		//获取用户提交数据
		if err := a.ParseForm(form); err == nil {
			user := models.GetUserByName(form.Name)
			if user == nil {
				//用户不存在
				errs.AddError("default", "用户名或密码错误")
			} else if user.ValidPassWord(form.Password) {
				//用户密码正确
				a.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
			} else {
				//用户密码不正确
				errs.AddError("default", "用户名或密码错误")
			}
		} else {
			errs.AddError("default", "用户名或密码错误")
		}

	}

	a.Data["form"] = form
	a.Data["errors"] = errs
	//定义加载界面
	a.TplName = "auth/login.html"
}
