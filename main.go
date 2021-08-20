package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "magego/course-33/cmdb/routers"
)

func main() {
	beego.Run()
}
