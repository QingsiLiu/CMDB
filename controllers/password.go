package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"html/template"
	"magego/course-33/cmdb/base/controllers/auth"
	"magego/course-33/cmdb/base/errors"
	"magego/course-33/cmdb/forms"
	"magego/course-33/cmdb/services"
	"net/http"
)

// PasswordController 用户修改密码控制器
type PasswordController struct {
	auth.LayoutController
}

// Modify 修改用户密码
func (p *PasswordController) Modify() {
	form := &forms.PasswordModifyForm{User: p.LoginUser}
	errs := errors.NewErrors()
	text := ""
	if p.Ctx.Input.IsPost() {
		if err := p.ParseForm(form); err == nil {
			//验证旧密码
			valid := &validation.Validation{}
			if success, err := valid.Valid(form); err != nil {
				errs.AddError("default", err.Error())
			} else if !success {
				errs.AddValidation(valid)
			} else {
				services.UserService.ModifyUserPassword(p.LoginUser.ID, form.Password)
				text = "修改密码成功"
				p.Redirect(beego.URLFor("AuthController.Logout"), http.StatusFound)
			}
			// 验证
			/*passwordRegexp := "^[0-9a-zA-Z_.\\$\\!#%^&\\*\\(\\)\\+]{6,20}$"
			valid.Match(form.Password, regexp.MustCompile(passwordRegexp), "default.default.default").Message("密码格式不正确")
			if valid.HasErrors() {
				for key, vErrors := range valid.ErrorsMap {
					for _, err := range vErrors {
						errs.AddError(key, err.Message)
					}
				}
			} else if form.Password != form.Password2 {
				errs.AddError("default", "两次密码不一致")
			} else if form.OldPassword == form.Password {
				errs.AddError("default", "新旧密码不能一致")
			} else {
				models.ModifyUserPassword(p.LoginUser.ID ,form.Password)
				text = "修改密码成功"
			}

			//验证密码范围:数字，大小写英文字母，特殊字符（_.$!#%^&*()+）,长度最小为6位，最长为20位
			passwordRegexp := "^[0-9a-zA-Z_.\\$\\!#%^&\\*\\(\\)\\+]{6,20}$"
			if isMatch, _ := regexp.MatchString(passwordRegexp, form.Password); !isMatch {
				errs.AddError("default", "密码只能由6~20位大小写字母，数字，特殊字符组成（_.$!#%^&*()+）")
			} else */
		}
	}
	p.TplName = "password/modify.html"
	p.Data["errors"] = errs
	p.Data["text"] = text
	p.Data["xsrf_input"] = template.HTML(p.XSRFFormHTML())
	p.Data["title"] = "修改密码"
}
