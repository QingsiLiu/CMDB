package models

import (
	gosql "database/sql"
	"magego/course-33/CMDB/utils"
	"time"
)

const (
	sqlQueryByName = "select id, name, password from user where name=?"
	sqlQuery       = "select id, staff_id, name, nickname, password, gender, tel, email, addr, department, status, created_at, updated_at, deleted_at from user"
)

type User struct {
	ID         int
	StaffID    string
	Name       string
	Nickname   string
	Password   string
	Gender     int
	Tel        string
	Addr       string
	Email      string
	Department string
	Status     int
	CreateAt   *time.Time
	UpdateAt   *time.Time
	DeleteAt   *time.Time
}

// GetUserByName 通过用户名获取用户
func GetUserByName(name string) *User {
	user := &User{}
	if err := db.QueryRow(sqlQueryByName, name).Scan(&user.ID, &user.Name, &user.Password); err == nil {
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
	users := make([]*User, 0, 10)
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
	}
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
