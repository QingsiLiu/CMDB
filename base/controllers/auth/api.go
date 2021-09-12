package auth

import (
	"fmt"
	"github.com/astaxie/beego"
	"magego/course-33/cmdb/base/controllers/base"
	"magego/course-33/cmdb/base/response"
)

type APIController struct {
	base.BaseController
}

func (c *APIController) Prepare() {
	c.EnableXSRF = false //针对Controller关闭xsrf检查

	token := fmt.Sprintf("Token %s", beego.AppConfig.DefaultString("api::token", ""))
	headerToken := c.Ctx.Input.Header("Authorization")

	if token != headerToken {
		c.Data["json"] = response.Unauthorization
	}
}

func (c *APIController) Render() error {
	c.ServeJSON()
	return nil
}
