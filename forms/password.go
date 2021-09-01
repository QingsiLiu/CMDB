package forms

import (
	"github.com/astaxie/beego/validation"
	"magego/course-33/cmdb/models"
	"regexp"
)

type PasswordModifyForm struct {
	User        *models.User `form:"-"`
	OldPassword string       `form:"old_password"`
	Password    string       `form:"password"`
	Password2   string       `form:"password2"`
}

// Valid
func (p *PasswordModifyForm) Valid(validation *validation.Validation) {
	if ok := p.User.ValidPassWord(p.OldPassword); !ok {
		validation.AddError("default.default", "旧密码错误")
		return
	}

	passwordRegexp := "^[0-9a-zA-Z_.\\$\\!#%^&\\*\\(\\)\\+]{6,20}$"
	validation.Match(p.Password, regexp.MustCompile(passwordRegexp), "default.default.default").Message("密码格式不正确")
	if validation.HasErrors() {
		return
	} else if p.Password != p.Password2 {
		validation.AddError("default.default", "两次密码不一致")
	} else if p.OldPassword == p.Password {
		validation.AddError("default.default", "新旧密码不能一致")
	}
}
