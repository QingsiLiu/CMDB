package services

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"magego/course-33/cmdb/forms"
	"magego/course-33/cmdb/models"
	"magego/course-33/cmdb/utils"
	"time"
)

type nodeService struct {
}

func (s *nodeService) Register(form *forms.NodeRegisterForm) {
	node := &models.Node{UUID: form.UUID}
	ormer := orm.NewOrm()
	if err := ormer.Read(node, "UUID"); err == nil {
		// 有数据更新
		node.HostName = form.HostName
		node.Addr = form.Addr
		node.DeletedAt = nil
		ormer.Update(node)
	} else if err == orm.ErrNoRows {
		// 没有数据就创建
		node.HostName = form.HostName
		node.Addr = form.Addr
		ormer.Insert(node)
	}
}

func (s *nodeService) GetByID(pk int) *models.Node {
	node := &models.Node{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(node, "id"); err == nil {
		return node
	}
	return nil
}

func (s *nodeService) Query(q string) []*models.Node {
	var nodes []*models.Node
	queryset := orm.NewOrm().QueryTable(&models.Node{})
	cond := orm.NewCondition()
	cond = cond.And("DeletedAt__isnull", true)

	if q != "" {
		cond1 := orm.NewCondition()
		cond1 = cond1.Or("HostName__icontains", q)
		cond1 = cond1.Or("Addr__icontains", q)
		cond = cond.AndCond(cond1)
	}

	queryset = queryset.SetCond(cond)
	queryset.All(&nodes)
	return nodes
}

func (s *nodeService) Delete(pk int) {
	if node := s.GetByID(pk); node != nil {
		deleteAt := time.Now()
		node.DeletedAt = &deleteAt
		ormer := orm.NewOrm()
		ormer.Update(node, "DeletedAt")
	}
}

//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------

type jobService struct {
}

func (s *jobService) GetByID(pk int) *models.Job {
	job := &models.Job{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(job, "id"); err == nil {
		return job
	}
	return nil
}

func (s *jobService) New(form *forms.JobForm) {
	job := &models.Job{
		Key:    form.Key,
		Remark: form.Remark,
		Node:   NodeService.GetByID(form.Node),
	}
	ormer := orm.NewOrm()
	ormer.Insert(job)
}

func (s *jobService) Modify(form *forms.JobForm) {
	if job := s.GetByID(form.ID); job != nil {
		job.Key = form.Key
		job.Remark = form.Remark
		job.Node = NodeService.GetByID(form.Node)
		ormer := orm.NewOrm()
		_, err := ormer.Update(job, "key", "Remark", "Node")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (s *jobService) Query(q string) []*models.Job {
	var jobs []*models.Job
	queryset := orm.NewOrm().QueryTable(&models.Job{})
	cond := orm.NewCondition()
	cond = cond.And("DeletedAt__isnull", true)

	if q != "" {
		cond1 := orm.NewCondition()
		cond1 = cond1.Or("Key__icontains", q)
		cond1 = cond1.Or("Remark__icontains", q)
		cond1 = cond1.Or("Node__HostName__icontains", q)
		cond1 = cond1.Or("Node__Addr__icontains", q)
		cond = cond.AndCond(cond1)
	}

	queryset = queryset.RelatedSel().SetCond(cond)
	queryset.All(&jobs)
	return jobs
}

func (s *jobService) QueryByUUID(uuid string) []*models.Job {
	var jobs []*models.Job
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&models.Job{})
	queryset.Filter("deleted_at__isnull", true).Filter("node__uuid", uuid).All(&jobs)
	for _, job := range jobs {
		ormer.LoadRelated(job, "Target")
	}
	return jobs
}

func (s *jobService) Delete(pk int) {
	if job := s.GetByID(pk); job != nil {
		deleteAt := time.Now()
		job.DeletedAt = &deleteAt
		ormer := orm.NewOrm()
		ormer.Update(job, "DeletedAt")
	}
}

//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------

type targetService struct {
}

func (s *targetService) GetByID(pk int) *models.Target {
	target := &models.Target{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(target, "id"); err == nil {
		return target
	}
	return nil
}

func (s *targetService) New(form *forms.TargetForm) *models.Target {
	target := &models.Target{
		Name:   form.Name,
		Remark: form.Remark,
		Addr:   form.Addr,
		Job:    JobService.GetByID(form.Job),
	}
	if _, err := orm.NewOrm().Insert(target); err == nil {
		return target
	}
	return nil
}

func (s *targetService) Delete(pk int) {
	if target := s.GetByID(pk); target != nil {
		deleteAt := time.Now()
		target.DeletedAt = &deleteAt
		ormer := orm.NewOrm()
		ormer.Update(target, "DeletedAt")
	}
}

func (s *targetService) Modify(form *forms.TargetForm) *models.Target {
	if target := s.GetByID(form.ID); target != nil {
		target.Name = form.Name
		target.Remark = form.Remark
		target.Addr = form.Addr
		target.Job = JobService.GetByID(form.Job)
		if _, err := orm.NewOrm().Update(target, "Name", "Remark", "Addr", "Job"); err == nil {
			return target
		}
	}
	return nil
}

func (s *targetService) Query(q string) []*models.Target {
	var targets []*models.Target
	queryset := orm.NewOrm().QueryTable(&models.Target{})
	cond := orm.NewCondition()
	cond = cond.And("DeletedAt__isnull", true)

	if q != "" {
		cond1 := orm.NewCondition()
		cond1 = cond1.Or("Name__icontains", q)
		cond1 = cond1.Or("Remark__icontains", q)
		cond1 = cond1.Or("Addr__icontains", q)
		cond = cond.AndCond(cond1)
	}

	queryset = queryset.RelatedSel().SetCond(cond)
	queryset.All(&targets)
	return targets
}

//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------

type alertService struct {
}

func (m *alertService) Alert(form *forms.AlertForm) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&models.Alert{})
	queryset = queryset.Filter("fingerprint", form.Fingerprint)
	queryset = queryset.Filter("deleted_at__isnull", true)
	queryset = queryset.Filter("status", "firing")
	if form.IsNew() {
		// 如果有为处理的告警，不在添加
		if count, err := queryset.Count(); err == nil && count == 0 {
			// 添加
			alert := &models.Alert{
				Fingerprint:  form.Fingerprint,
				Alertname:    form.AlertName(),
				Status:       form.Status,
				StartsAt:     form.StartsAt,
				GeneratorURL: form.GeneratorURL,
				Labels:       form.LabelsString(),
				Annotations:  form.AnnotationsString(),
			}
			ormer.Insert(alert)
		}
	} else {
		// 更新
		queryset.Update(orm.Params{
			"EndsAt": form.EndsAt,
			"Status": form.Status,
		})
	}
}

// Query 查询
func (s *alertService) Query(form *forms.AlertQueryParams) *utils.Page {
	var alerts []*models.Alert
	queryset := orm.NewOrm().QueryTable(&models.Alert{})
	cond := orm.NewCondition()
	cond = cond.And("deleted_at__isnull", true)

	if form.Q != "" {
		qcond := orm.NewCondition()
		qcond = qcond.Or("alertname__icontains", form.Q)
		cond = cond.AndCond(qcond)
	}

	if form.Status != "" && form.Status != "all" {
		qcond := orm.NewCondition()
		qcond = qcond.Or("status", form.Status)
		cond = cond.AndCond(qcond)
	}

	if form.StartTime() != nil {
		qcond := orm.NewCondition()
		qcond = qcond.Or("created_at__gte", form.StartTime())
		cond = cond.AndCond(qcond)
	}

	if form.EndTime() != nil {
		qcond := orm.NewCondition()
		qcond = qcond.Or("created_at__lt", form.EndTime())
		cond = cond.AndCond(qcond)
	}

	queryset.SetCond(cond).OrderBy("-created_at").Offset(form.Offset()).Limit(form.PageSize()).All(&alerts)
	total, _ := queryset.SetCond(cond).Count()
	// NewPage(alerts, total, form.PageQueryParams)
	return utils.NewPage(alerts, total, form.PageSize(), form.PageNum(), form.Inputs)
}

func (s *alertService) Notice(form *forms.AlertGroupForm) {
	// 告警永远通知个某个，某些人，某个群组 => 通知所有运维人员
	// 业务发生故障 => 通知业务负责人
	// 告警分组 => 业务 ==> cmdb 业务 => 负责人 (告警规则)
	tos := beego.AppConfig.DefaultStrings("notice::mailTos", []string{})
	subject := form.AlertName()

	phones := beego.AppConfig.DefaultStrings("notice::phones", []string{})
	templateId := beego.AppConfig.DefaultString("notice::templateId", "")
	templateParams := []string{form.AlertName(), "账户余额", "CMDB"}

	utils.SendMail(tos, subject, utils.FormatEmailBody("views/email/alert.html", form))
	utils.SendSms(phones, templateId, templateParams)
}

var (
	NodeService   = new(nodeService)
	JobService    = new(jobService)
	TargetService = new(targetService)
	AlertService  = new(alertService)
)
