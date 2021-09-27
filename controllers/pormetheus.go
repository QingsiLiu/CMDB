package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"magego/course-33/cmdb/base/controllers/auth"
	"magego/course-33/cmdb/forms"
	"magego/course-33/cmdb/services"
	"net/http"
)

type prometheusController struct {
	auth.LayoutController
}

func (p *prometheusController) Prepare() {
	p.LayoutController.Prepare()
	p.Data["nav"] = "prometheus"
	p.Data["subnav"] = p.GetNav()
}

//-------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------

type NodeController struct {
	prometheusController
}

// Query 查询节点
func (n *NodeController) Query() {
	q := n.GetString("q")

	n.Data["nodes"] = services.NodeService.Query(q)
	n.Data["q"] = q
	n.TplName = "prometheus/node/query.html"
	n.Data["title"] = "节点列表"
}

// Delete 删除节点
func (c *NodeController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil {
		services.NodeService.Delete(pk)
	}

	c.Redirect(beego.URLFor("NodeController.Query"), http.StatusFound)
}

//-------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------

type JobController struct {
	prometheusController
}

// New 新增任务
func (c *JobController) New() {
	form := &forms.JobForm{}

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			services.JobService.New(form)
			c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
		}
	}

	c.Data["form"] = form
	c.Data["title"] = "任务新建"
	c.Data["nodes"] = services.NodeService.Query("")
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "prometheus/job/new.html"
}

// Modify 编辑任务
func (c *JobController) Modify() {
	form := &forms.JobForm{}

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证数据
			services.JobService.Modify(form)
			c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
		}
	} else if pk, err := c.GetInt("pk"); err == nil {
		if job := services.JobService.GetByID(pk); job != nil {
			form.ID = job.ID
			form.Key = job.Key
			form.Remark = job.Remark
			form.Node = job.Node.ID
			fmt.Println(form)
		}
	}

	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["nodes"] = services.NodeService.Query("")
	c.Data["title"] = "用户修改"
	c.TplName = "prometheus/job/modify.html"
}

// Query 查询任务
func (c *JobController) Query() {
	beego.ReadFromRequest(&c.Controller)

	q := c.GetString("q")

	c.Data["jobs"] = services.JobService.Query(q)
	c.Data["q"] = q
	c.TplName = "prometheus/job/query.html"
	c.Data["title"] = "任务列表"
}

// Delete 删除任务
func (c *JobController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil {
		services.JobService.Delete(pk)
	}

	c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
}

//-------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------

type TargetController struct {
	prometheusController
}

// New 新增目标
func (c *TargetController) New() {
	form := &forms.TargetForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			services.TargetService.New(form)
			c.Redirect(beego.URLFor("TargetController.Query"), http.StatusFound)
		}
	}

	c.Data["form"] = form
	c.Data["title"] = "目标新建"
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["jobs"] = services.JobService.Query("")
	c.TplName = "prometheus/target/new.html"
}

// Modify 编辑目标
func (c *TargetController) Modify() {
	form := &forms.TargetForm{}

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			// 验证数据
			services.TargetService.Modify(form)
			c.Redirect(beego.URLFor("TargetController.Query"), http.StatusFound)
		}
	} else if pk, err := c.GetInt("pk"); err == nil {
		if target := services.TargetService.GetByID(pk); target != nil {
			form.ID = target.ID
			form.Name = target.Name
			form.Addr = target.Addr
			form.Remark = target.Remark
			form.Job = target.Job.ID
		}
	}

	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["title"] = "目标修改"
	c.Data["jobs"] = services.JobService.Query("")
	c.TplName = "prometheus/target/modify.html"
}

// Delete 删除任务
func (c *TargetController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil {
		services.TargetService.Delete(pk)
	}

	c.Redirect(beego.URLFor("TargetController.Query"), http.StatusFound)
}

// Query 查询目标
func (t *TargetController) Query() {
	beego.ReadFromRequest(&t.Controller)

	q := t.GetString("q")

	t.Data["targets"] = services.TargetService.Query(q)
	t.Data["q"] = q
	t.TplName = "prometheus/target/query.html"
	t.Data["title"] = "目标列表"
}

//-------------------------------------------------------------------------------------
//-------------------------------------------------------------------------------------
type AlertController struct {
	prometheusController
}

func (c *AlertController) Query() {
	form := forms.NewAlertQueryParams(c.Input())
	if err := c.ParseForm(form); err == nil {
		fmt.Printf("%#v\n", form.PageQueryParams)
		fmt.Printf("%#v\n", form)
		c.Data["page"] = services.AlertService.Query(form)
	}
	c.Data["form"] = form
	c.TplName = "prometheus/alert/query.html"
}
