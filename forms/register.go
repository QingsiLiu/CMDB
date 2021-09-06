package forms

import (
	"github.com/astaxie/beego/validation"
	"regexp"
)

type RegisterForm struct {
	Name      string `form:"name"`
	Password  string `form:"password"`
	Password2 string `form:"password2"`
}

// Valid 注册密码的验证
func (r *RegisterForm) Valid(validation *validation.Validation) {
	passwordRegexp := "^[0-9a-zA-Z_.\\$\\!#%^&\\*\\(\\)\\+]{6,20}$"
	validation.Match(r.Password, regexp.MustCompile(passwordRegexp), "default.default.default").Message("密码格式不正确")
	if validation.HasErrors() {
		return
	} else if r.Password != r.Password2 {
		validation.AddError("default.default", "两次密码不一致")
	}
}
