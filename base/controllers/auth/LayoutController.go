package auth

type LayoutController struct {
	AuthorizationController
}

func (l *LayoutController) Prepare() {
	l.AuthorizationController.Prepare()

	l.LayoutSections = make(map[string]string)
	l.LayoutSections["SectionStyle"] = ""

	l.Layout = "base/layouts/layout.html"
}
