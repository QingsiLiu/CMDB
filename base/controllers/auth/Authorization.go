package auth

import (
	"github.com/astaxie/beego"
	"magego/course-33/cmdb/base/controllers/base"
	"magego/course-33/cmdb/models"
	"net/http"
	"strings"
)

// AuthorizationController 所有需要认证访问基础控制器
type AuthorizationController struct {
	base.BaseController
	LoginUser *models.User
}

func (a *AuthorizationController) getNav() string {
	controllerName, _ := a.GetControllerAndAction()
	return strings.ToLower(strings.TrimSuffix(controllerName, "Controller"))
}

// Prepare 用户认证检查
func (a *AuthorizationController) Prepare() {
	//配置文件读取
	SessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	SessionUser := a.GetSession(SessionKey)
	a.Data["LoginUser"] = nil
	a.Data["nav"] = a.getNav()

	if SessionUser != nil {
		if id, ok := SessionUser.(int); ok {
			if user := models.GetUserByID(id); user != nil {
				a.Data["loginUser"] = user
				a.LoginUser = user
				return
			}
		}
	}

	LoginController := beego.AppConfig.DefaultString("auth::LoginController", "AuthController.Login")
	a.Redirect(beego.URLFor(LoginController), http.StatusFound)

}
