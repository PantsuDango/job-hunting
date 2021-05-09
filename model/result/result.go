package result

import (
	"job-hunting/model/params"
)

type ListJob struct {
	Count   int                `json:"Count"`
	ListJob []params.JobParams `json:"ListJob"`
}

type UserInfo struct {
	DeliverCount int  `json:"DeliverCount"`
	UserInfo     User `json:"UserInfo"`
}

type User struct {
	ID        int    `json:"ID"`
	Nick      string `json:"Nick"`
	UserName  string `json:"UserName"`
	Sex       int    `json:"Sex"`
	HeadImage string `json:"HeadImage"`
	Email     string `json:"Email"`
	Phone     string `json:"Phone"`
	Birthday  string `json:"Birthday"`
	Degree    string `json:"Degree"`
	Job       string `json:"Job"`
	Address   string `json:"Address"`
}
