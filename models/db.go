package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"log"
)

var db *sql.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/cmdb?charset=utf8mb4&loc=PRC&parseTime=true",
		beego.AppConfig.DefaultString("mysql::User", "root"),
		beego.AppConfig.DefaultString("mysql::Password", "root"),
		beego.AppConfig.DefaultString("mysql::Host", "127.0.0.1"),
		beego.AppConfig.DefaultInt("ysql::Port", 3306))

	db, _ = sql.Open("mysql", dsn)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {

	}
}
