package models

import (
	"magego/course-33/cmdb/utils"
)

const (
	sqlQueryByName = "select id, name, password from user where name=?"
	sqlQuery       = "select id, name from user"
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
	return u.Password == utils.Md5Text(password)
}

// QueryUser 查询用户
func QueryUser(q string) []*User {
	users := make([]*User, 0, 10)
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil
	}
	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.ID, &user.Name); err == nil {
			users = append(users, user)
		}
	}
	return users
}
