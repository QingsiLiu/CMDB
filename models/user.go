package models

import (
	"github.com/astaxie/beego/orm"
	"magego/course-33/cmdb/forms"
	"magego/course-33/cmdb/utils"
	"time"
)

type User struct {
	ID         int        `orm:"column(id)"`
	StaffID    string     `orm:"column(staff_id);size(32)"`
	Name       string     `orm:"size(64)"`
	Nickname   string     `orm:"size(64)"`
	Password   string     `orm:"size(1024)"`
	Gender     int        `orm:""`
	Tel        string     `orm:"size(32)"`
	Addr       string     `orm:"size(128)"`
	Email      string     `orm:"size(64)"`
	Department string     `orm:"size(128)"`
	Status     int        `orm:""`
	CreatedAt  *time.Time `orm:"auto_now_add"`
	UpdatedAt  *time.Time `orm:"auto_now"`
	DeletedAt  *time.Time `orm:"null"`
}

func init() {
	orm.RegisterModel(new(User))
}

// GetUserByName 通过用户名获取用户
func GetUserByName(name string) *User {
	user := &User{Name: name}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

// GetUserByID 通过ID获取用户
func GetUserByID(id int) *User {
	user := &User{ID: id}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "id"); err == nil {
		return user
	}
	return nil
}

// ValidPassWord 验证用户密码是否正确
func (u *User) ValidPassWord(password string) bool {
	return utils.CheckPassword(password, u.Password)
}

// QueryUser 查询用户
func QueryUser(q string) []*User {
	var users []*User
	queryset := orm.NewOrm().QueryTable(&User{})
	cond := orm.NewCondition()
	cond = cond.And("DeletedAt__isnull", true)

	if q != "" {
		cond1 := orm.NewCondition()
		cond1 = cond1.Or("Name__icontains", q)
		cond1 = cond1.Or("Nickname__icontains", q)
		cond1 = cond1.Or("Tel__icontains", q)
		cond1 = cond1.Or("Addr__icontains", q)
		cond1 = cond1.Or("Email__icontains", q)
		cond1 = cond1.Or("Department__icontains", q)
		cond = cond.AndCond(cond1)
	}

	queryset = queryset.SetCond(cond)
	queryset.All(&users)
	return users
}

// NewUser 新建用户信息
func NewUser(form *forms.UserModifyForm) {
	user := &User{StaffID: form.StaffId,
		Name:       form.Name,
		Nickname:   form.NickName,
		Gender:     form.GenderInt(),
		Department: form.Department,
		Tel:        form.Tel,
		Addr:       form.Addr,
		Email:      form.Email,
		Status:     form.StatusInt()}
	ormer := orm.NewOrm()
	ormer.Insert(user)
}

// ModifyUser 修改用户信息
func ModifyUser(form *forms.UserModifyForm) {
	if user := GetUserByID(form.ID); user != nil {
		user.StaffID = form.StaffId
		user.Nickname = form.NickName
		user.Password = utils.GeneratePassword(form.Password)
		user.Gender = form.GenderInt()
		user.Tel = form.Tel
		user.Email = form.Email
		user.Addr = form.Addr
		user.Department = form.Department
		user.Status = form.StatusInt()
		ormer := orm.NewOrm()
		ormer.Update(user, "StaffID", "Nickname", "Password", "Gender", "Tel", "Email", "Addr", "Department", "Status")
	}
}

// DeleteUser 删除用户信息
func DeleteUser(pk int) {
	/*ormer := orm.NewOrm()
	ormer.Delete(&User{ID: pk})*/
	if user := GetUserByID(pk); user != nil {
		deleteAt := time.Now()
		user.DeletedAt = &deleteAt
		ormer := orm.NewOrm()
		ormer.Update(user, "DeletedAt")
	}
}

func (u *User) GenderText() string {
	if u.Gender == 0 {
		return "女"
	} else if u.Gender == 1 {
		return "男"
	} else {
		return "未知"
	}
}

func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "正常"
	case 1:
		return "锁定"
	case 2:
		return "离职"
	}
	return "未知"
}
