package controllers

import (
	"github.com/astaxie/beego"
	"magego/course-33/cmdb/base/controllers/auth"
	"magego/course-33/cmdb/forms"
	"magego/course-33/cmdb/services"
	"magego/course-33/cmdb/utils"
	"net/http"
	"strings"
	"time"
)

const TimeLayout = "2006-01-02 15:04"

type TaskController struct {
	auth.LayoutController
}

// formatTime 时间格式化
func formatTime(t string) *time.Time {
	if t != "" {
		ft, _ := time.Parse(TimeLayout, strings.ReplaceAll(t, "T", " "))
		return &ft
	}
	return nil
}

// New 添加任务
func (t *TaskController) New() {
	taskform := &forms.TaskForm{}

	if t.Ctx.Input.IsPost() {
		if err := t.ParseForm(taskform); err == nil {
			services.TaskService.New(taskform)
			t.Redirect(beego.URLFor("TaskController.Query"), http.StatusFound)
		}
	}

	t.Data["xsrf_token"] = t.XSRFToken()
	t.TplName = "task/new.html"
	t.Data["title"] = "任务新建"
}

// Delete 删除任务列表
func (t *TaskController) Delete() {
	if pk, err := t.GetInt("pk"); err == nil {
		services.TaskService.Delete(pk)
	}

	t.Redirect(beego.URLFor("TaskController.Query"), http.StatusFound)
}

// Modify 修改任务列表
func (t *TaskController) Modify() {
	taskForm := &forms.TaskForm{}

	if t.Ctx.Input.IsPost() {
		if err := t.ParseForm(taskForm); err == nil {
			services.TaskService.Modify(taskForm)

			// 存储消息
			flash := beego.NewFlash()
			flash.Set("notice", "修改任务信息成功")
			flash.Store(&t.Controller)

			t.Redirect(beego.URLFor("TaskController.Query"), http.StatusFound)
		}
	} else if pk, err := t.GetInt("pk"); err == nil {
		if task := services.TaskService.GetTaskById(pk); task != nil {
			taskForm.ID = task.ID
			taskForm.Name = task.Name
			taskForm.StartTime = utils.Time2String(task.StartTime)
			taskForm.CompleteTime = utils.Time2String(task.CompleteTime)
			taskForm.DeadlineTime = utils.Time2String(task.DeadlineTime)
			taskForm.Content = task.Content
		}
	}

	t.Data["form"] = taskForm
	t.Data["xsrf_token"] = t.XSRFToken()
	t.TplName = "task/modify.html"
	t.Data["title"] = "任务修改"
}

// Query 查询任务列表
func (t *TaskController) Query() {
	beego.ReadFromRequest(&t.Controller)

	q := t.GetString("q")
	t.Data["tasks"] = services.TaskService.Query(q)
	t.TplName = "task/query.html"
	t.Data["title"] = "任务列表"
}
