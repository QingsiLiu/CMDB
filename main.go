package main

import (
	_ "github.com/go-sql-driver/mysql"
	"magego/course-33/cmdb/cmds"
	_ "magego/course-33/cmdb/routers"
)

func main() {
	cmds.Execute()
}
