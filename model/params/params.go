package params

import "time"

type ModActIndex struct {
	Module string `json:"Module" form:"Module" binding:"required"`
	Action string `json:"Action" form:"Action" binding:"required"`
}

type JobParams struct {
	Id          int       `json:"Id"`
	Name        string    `json:"Name"         binding:"required"`
	Pay         string    `json:"Pay"          binding:"required"`
	IcoUrl      string    `json:"IcoUrl"`
	Company     string    `json:"Company"      binding:"required"`
	Scale       string    `json:"Scale"`
	Description string    `json:"Description"`
	Tags        []string  `json:"Tags"`
	CreatedAt   time.Time `json:"-"            gorm:"column:createtime"`
	Createtime  string    `json:"Createtime"`
}

func (JobParams) TableName() string {
	return "job"
}

type ListJobParams struct {
	Offset int `json:"Offset"`
	Limit  int `json:"Limit"`
}
