package forms

// UserModifyForm 用户修改表单
type UserModifyForm struct {
	ID         int    `form:"id"`
	StaffId    string `form:"staffId"`
	Name       string `form:"name"`
	NickName   string `form:"nickname"`
	Gender     string `form:"gender"`
	Tel        string `form:"tel"`
	Email      string `form:"email"`
	Addr       string `form:"addr"`
	Department string `form:"department"`
	Status     string `form:"status"`
}

func (u *UserModifyForm) GenderInt() int {
	if u.Gender == "女" {
		return 0
	} else if u.Gender == "男" {
		return 1
	} else {
		return 2
	}
}

func (u *UserModifyForm) StatusInt() int {
	switch u.Status {
	case "正常":
		return 0
	case "锁定":
		return 1
	case "离职":
		return 2
	}
	return 3
}
