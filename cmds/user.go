package cmds

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cobra"
	"magego/course-33/cmdb/models"
	"magego/course-33/cmdb/utils"
)

var (
	name     string
	password string
)

var userCommand = &cobra.Command{
	Use:   "user",
	Short: "user console",
	Long:  "user console",
	RunE: func(cmd *cobra.Command, args []string) error {
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

		ormer := orm.NewOrm()
		user := &models.User{Name: name}
		user.Password = utils.GeneratePassword(password)
		_, err := ormer.Insert(user)
		return err
	},
}

func init() {
	rootCommand.AddCommand(userCommand)
	dbCommand.Flags().StringVarP(&name, "name", "n", "admin", "name")
	dbCommand.Flags().StringVarP(&password, "password", "p", "root", "password")
}
