package result

import (
	"job-hunting/model/params"
)

type ListJob struct {
	Count   int                `json:"Count"`
	ListJob []params.JobParams `json:"ListJob"`
}

type UserInfo struct {
	DeliverCount   int            `json:"DeliverCount"`
	UserInfo       User           `json:"UserInfo"`
	ResumeInfo     ResumeInfo     `json:"ResumeInfo"`
	EducationInfo  EducationInfo  `json:"EducationInfo"`
	JobExpectation JobExpectation `json:"JobExpectation"`
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

type ResumeInfo struct {
	State          string `json:"State"`
	City           string `json:"City"`
	Identity       string `json:"Identity"`
	Intention      string `json:"Intention"`
	Advantage      string `json:"Advantage"`
	WorkExperience string `json:"WorkExperience"`
}

type EducationInfo struct {
	SchoolName        string `json:"SchoolName"`
	Major             string `json:"Major"`
	MatriculationTime string `json:"MatriculationTime"`
	GraduationTime    string `json:"GraduationTime"`
}

type JobExpectation struct {
	JobTags string `json:"JobTags"`
	Pay     string `json:"Pay"`
	City    string `json:"City"`
}
