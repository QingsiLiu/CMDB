package auth

import "strings"

type LayoutController struct {
	AuthorizationController
}

// GetNav 获取菜单active
func (l *LayoutController) GetNav() string {
	controllerName, _ := l.GetControllerAndAction()
	return strings.ToLower(strings.TrimSuffix(controllerName, "Controller"))
}

func (l *LayoutController) Prepare() {
	l.AuthorizationController.Prepare()

	l.LayoutSections = make(map[string]string)
	l.LayoutSections["SectionStyle"] = ""

	l.Data["nav"] = l.GetNav()
	l.Data["subnav"] = ""

	l.Layout = "base/layouts/layout.html"
}
