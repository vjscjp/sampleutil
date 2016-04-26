package controllers

type App struct {
	Id          string  `json:"id"`
	ProjectName string  `json:"projectname"`
	ProjectId   string  `json:"projectid"`
	EnvName     string  `json:"envname"`
	EnvId       string  `json:"envid"`
	ServiceName string  `json:"servicename"`
	ServiceId   string  `json:"serviceid"`
	Hosts       []*Host `json:"hosts,omitempty"`
}

type Host struct {
	Host  string `json:"host"`
	Ports []int  `json:"ports"`
}
