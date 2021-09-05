package controllers

import (
	"magego/course-33/cmdb/base/controllers/auth"
)

type HomeController struct {
	auth.LayoutController
}

func (h *HomeController) Index() {
	h.TplName = "home/index.html"
	h.Data["title"] = "首页"
}
