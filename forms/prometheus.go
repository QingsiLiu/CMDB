package forms

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
