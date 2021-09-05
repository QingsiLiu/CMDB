package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"magego/course-33/cmdb/config"
	_ "magego/course-33/cmdb/routers"
)

func main() {
	beego.SetLogger("file", `{"filename" : "logs/cmdb.log"}`)
	beego.SetLogFuncCall(true)
	beego.SetLevel(beego.LevelDebug)
	beego.BeeLogger.DelLogger("console")

	config.Init("file", `{"CachePath" : "cache", "FileSuffix" : ".cache", "EmbedExpiry" : "60", "Directory" : "3"}`)

	orm.Debug = true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/cmdb?charset=utf8mb4&loc=PRC&parseTime=true",
		beego.AppConfig.DefaultString("mysql::User", "root"),
		beego.AppConfig.DefaultString("mysql::Password", "root"),
		beego.AppConfig.DefaultString("mysql::Host", "127.0.0.1"),
		beego.AppConfig.DefaultInt("mysql::Port", 3306))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)

	if db, err := orm.GetDB(); err != nil {
		log.Fatal(err)
	} else if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	beego.Run()
}
