package controllers

import (
	"fmt"
	"magego/course-33/cmdb/base/controllers/auth"
	"magego/course-33/cmdb/config"
	"time"
)

type HomeController struct {
	auth.LayoutController
}

func (h *HomeController) Index() {
	v := config.Cache.Get("stime")
	if v != nil {
		vv, ok := v.([]byte)
		fmt.Println(string(vv), ok)
	}
	config.Cache.Put("stime", time.Now().Format("2006-01-02 15:04:05"), time.Minute)

	h.TplName = "home/index.html"
	h.Data["title"] = "首页"
}
