package services

import (
	"github.com/astaxie/beego/orm"
	"magego/course-33/cmdb/forms"
	"magego/course-33/cmdb/models"
	"magego/course-33/cmdb/utils"
	"time"
)

type userService struct {
}

// GetByName 通过用户名获取用户
func (u *userService) GetByName(name string) *models.User {
	user := &models.User{Name: name}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

// GetByID 通过ID获取用户
func (u *userService) GetByID(id int) *models.User {
	user := &models.User{ID: id}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "id"); err == nil {
		return user
	}
	return nil
}

func (u *userService) Register(form *forms.RegisterForm) {
	user := &models.User{
		Name:     form.Name,
		Password: utils.GeneratePassword(form.Password),
		Gender:   3,
		Status:   4,
	}
	ormer := orm.NewOrm()
	ormer.Insert(user)
}

// New 新建用户信息
func (u *userService) New(form *forms.UserModifyForm) {
	user := &models.User{StaffID: form.StaffId,
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

// Query 查询用户
func (u *userService) Query(q string) []*models.User {
	var users []*models.User
	queryset := orm.NewOrm().QueryTable(&models.User{})
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

// Modify 修改用户信息
func (u *userService) Modify(form *forms.UserModifyForm) {
	if user := u.GetByID(form.ID); user != nil {
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

// Delete 删除用户信息
func (u *userService) Delete(pk int) {
	/*ormer := orm.NewOrm()
	ormer.Delete(&User{ID: pk})*/
	if user := u.GetByID(pk); user != nil {
		deleteAt := time.Now()
		user.DeletedAt = &deleteAt
		ormer := orm.NewOrm()
		ormer.Update(user, "DeletedAt")
	}
}

// ModifyUserPassword 修改用户密码
func (u *userService) ModifyUserPassword(pk int, password string) {
	if user := u.GetByID(pk); user != nil {
		user.Password = utils.GeneratePassword(password)
		ormer := orm.NewOrm()
		ormer.Update(user, "Password")
	}
}

// UserService 用户操作业务
var UserService = new(userService)
