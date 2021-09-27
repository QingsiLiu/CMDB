package forms

import (
	"encoding/json"
	"magego/course-33/cmdb/utils"
	"net/url"
	"time"
)

type NodeRegisterForm struct {
	UUID     string `json:"uuid"`
	HostName string `json:"hostname"`
	Addr     string `json:"addr"`
}

type JobForm struct {
	ID     int    `form:"ID"`
	Key    string `form:"Key"`
	Remark string `form:"Remark"`
	Node   int    `form:"Node"`
}

type TargetForm struct {
	ID     int    `form:"ID"`
	Name   string `form:"Name"`
	Remark string `form:"Remark"`
	Addr   string `form:"Addr"`
	Job    int    `form:"Job"`
}

type AlertForm struct {
	Fingerprint  string            `json:"fingerprint"`
	Status       string            `json:"status"`
	StartsAt     *time.Time        `json:"startsAt"`
	EndsAt       *time.Time        `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
}

func (f *AlertForm) IsNew() bool {
	return f.Status == "firing"
}

func (f *AlertForm) AlertName() string {
	return f.Labels["alertname"]
}

func (f *AlertForm) LabelsString() string {
	if bytes, err := json.Marshal(f.Labels); err == nil {
		return string(bytes)
	}
	return "{}"
}

func (f *AlertForm) AnnotationsString() string {
	if bytes, err := json.Marshal(f.Annotations); err == nil {
		return string(bytes)
	}
	return "{}"
}

type AlertQueryParams struct {
	utils.PageQueryParams

	Q      string `form:"q"`
	Status string `form:"status"`
	Stime  string `form:"stime"`
	Etime  string `form:"etime"`
}

func NewAlertQueryParams(inputs url.Values) *AlertQueryParams {
	return &AlertQueryParams{PageQueryParams: utils.PageQueryParams{Inputs: inputs}}
}

func (f *AlertQueryParams) StartTime() *time.Time {
	loc, _ := time.LoadLocation("PRC") // 写入到配置文件
	if t, err := time.ParseInLocation("2006-01-02T15:04", f.Stime, loc); err == nil {
		// if t, err := time.Parse("2006-01-02T15:04", f.Stime); err == nil {
		return &t
	}
	return nil
}

func (f *AlertQueryParams) EndTime() *time.Time {
	loc, _ := time.LoadLocation("PRC")
	if t, err := time.ParseInLocation("2006-01-02T15:04", f.Etime, loc); err == nil {
		// if t, err := time.Parse("2006-01-02T15:04", f.Etime); err == nil {
		return &t
	}
	return nil
}

type AlertGroupForm struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
	Status            string            `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Alerts            []*AlertForm      `json:"alerts"`
}

func (f *AlertGroupForm) AlertName() string {
	return f.GroupLabels["alertname"]
}
