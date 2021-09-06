package services

import (
	"github.com/astaxie/beego/orm"
	"magego/course-33/cmdb/forms"
	"magego/course-33/cmdb/models"
	"magego/course-33/cmdb/utils"
)

type taskService struct {
}

// TaskService 用户操作业务
var TaskService = new(taskService)

// New 任务新建操作
func (t *taskService) New(form *forms.TaskForm) {
	task := &models.Task{
		Name:         form.Name,
		Content:      form.Content,
		Status:       form.StatusInt(),
		User:         form.User,
		StartTime:    utils.String2Time(form.StartTime),
		DeadlineTime: utils.String2Time(form.DeadlineTime),
		CompleteTime: utils.String2Time(form.CompleteTime),
	}
	ormer := orm.NewOrm()
	ormer.Insert(&task)
}

// Query 任务查询操作
func (t *taskService) Query(q string) []*models.Task {
	var tasks []*models.Task
	queryset := orm.NewOrm().QueryTable(&models.Task{})
	cond := orm.NewCondition()
	cond = cond.And("DeletedAt__isnull", true)

	if q != "" {
		cond1 := orm.NewCondition()
		cond1 = cond1.Or("ID__icontains", q)
		cond1 = cond1.Or("Name__icontains", q)
		cond1 = cond1.Or("Content__icontains", q)
		cond = cond.AndCond(cond1)
	}

	queryset = queryset.SetCond(cond)
	queryset.All(&tasks)
	return tasks
}
