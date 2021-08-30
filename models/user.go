package models

import (
	"github.com/astaxie/beego/orm"
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
	if q != "" {
		cond := orm.NewCondition()
		cond = cond.Or("Name__icontains", q)
		cond = cond.Or("Nickname__icontains", q)
		cond = cond.Or("Tel__icontains", q)
		cond = cond.Or("Addr__icontains", q)
		cond = cond.Or("Email__icontains", q)
		cond = cond.Or("Department__icontains", q)
		queryset = queryset.SetCond(cond)
	}
	queryset.All(&users)
	return users
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
