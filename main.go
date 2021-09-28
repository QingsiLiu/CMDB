package main

import (
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	"magego/course-33/cmdb/cmds"
	_ "magego/course-33/cmdb/routers"
)

func main() {
	cmds.Execute()
}
