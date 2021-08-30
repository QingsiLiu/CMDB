package models

import (
	"github.com/astaxie/beego/orm"
	"magego/course-33/CMDB/utils"
	"time"
)

/*const (
	sqlQueryByName = "select id, name, password from user where name=?"
	sqlQuery       = "select id, staff_id, name, nickname, password, gender, tel, email, addr, department, status, created_at, updated_at, deleted_at from user"
)*/

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
	/*user := &User{}
	if err := db.QueryRow(sqlQueryByName, name).Scan(&user.ID, &user.Name, &user.Password); err == nil {
		return user
	}
	return nil*/
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
	/*users := make([]*User, 0, 10)
	sql := sqlQuery
	var (
		rows *gosql.Rows
		err  error
	)
	var params []interface{}
	q = utils.Like(q)
	if q != "" {
		sql += " WHERE staff_id like ? ESCAPE '/' OR name like ? ESCAPE '/' OR nickname like ? ESCAPE '/' OR tel like ? ESCAPE '/' OR email like ? ESCAPE '/' OR addr like ? ESCAPE '/' OR department like ? ESCAPE '/'"
		params = append(params, q, q, q, q, q, q, q)
	}

	rows, err = db.Query(sql, params...)
	if err != nil {
		return users
	}

	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.ID, &user.StaffID, &user.Name, &user.Nickname, &user.Password, &user.Gender,
			&user.Tel, &user.Email, &user.Addr, &user.Department, &user.Status, &user.CreateAt,
			&user.UpdateAt, &user.DeleteAt); err == nil {
			users = append(users, user)
		}
	}*/
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
