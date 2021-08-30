package auth

import (
	"github.com/astaxie/beego"
	"magego/course-33/cmdb/base/controllers/base"
	"net/http"
)

// AuthorizationController 所有需要认证访问基础控制器
type AuthorizationController struct {
	base.BaseController
}

// Prepare 用户认证检查
func (a *AuthorizationController) Prepare() {
	//配置文件读取
	SessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	LoginController := beego.AppConfig.DefaultString("auth::LoginController", "AuthController.Login")
	SessionUser := a.GetSession(SessionKey)
	if SessionUser == nil {
		a.Redirect(beego.URLFor(LoginController), http.StatusFound)
	}
	//查询用户信息 => Data
	//根据ID获取用户数据

}
