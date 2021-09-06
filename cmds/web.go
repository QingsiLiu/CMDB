package cmds

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cobra"
	"magego/course-33/cmdb/config"
)

var webCommand = &cobra.Command{
	Use:   "web",
	Short: "web console",
	Long:  "web console",
	RunE: func(cmd *cobra.Command, args []string) error {
		/*beego.SetLogger("file", `{"filename" : "logs/cmdb.log"}`)
		beego.SetLogFuncCall(true)
		beego.BeeLogger.DelLogger("console")*/

		config.Init("file", `{"CachePath" : "tmp/cache", "FileSuffix" : ".cache", "EmbedExpiry" : "60", "Directory" : "3"}`)

		orm.Debug = verbose
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/cmdb?charset=utf8mb4&loc=PRC&parseTime=true",
			beego.AppConfig.DefaultString("mysql::User", "root"),
			beego.AppConfig.DefaultString("mysql::Password", "root"),
			beego.AppConfig.DefaultString("mysql::Host", "127.0.0.1"),
			beego.AppConfig.DefaultInt("mysql::Port", 3306))
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", dsn)

		if db, err := orm.GetDB(); err != nil {
			return err
		} else if err := db.Ping(); err != nil {
			return err
		}

		beego.Run()
		return nil
	},
}

func init() {
	rootCommand.AddCommand(webCommand)
}
