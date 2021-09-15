package v1

import (
	"encoding/json"
	"fmt"
	"magego/course-33/cmdb/base/controllers/auth"
	"magego/course-33/cmdb/base/response"
	"magego/course-33/cmdb/forms"
	"magego/course-33/cmdb/services"
)

type PrometheusController struct {
	auth.APIController
}

func (c *PrometheusController) Register() {
	form := &forms.NodeRegisterForm{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, form); err == nil {
		//验证
		services.NodeService.Register(form)
		c.Data["json"] = response.OK
	} else {
		fmt.Println(err)
		c.Data["json"] = response.BadRequest
	}

}

func (c *PrometheusController) Config() {
	uuid := c.GetString("uuid")
	// job target
	/*
		[
			{
				"key" : " ",
				"targets": [
					{"addr", ""}, {"addr" : ""}
				]
			}
		]
	*/
	rt := services.JobService.QueryByUUID(uuid)

	c.Data["json"] = response.NewJSONResponse(200, "ok", rt)
}

func (c *PrometheusController) Alert() {
	fmt.Println(string(c.Ctx.Input.RequestBody))
}
