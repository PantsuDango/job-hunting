package tables

import "time"

type User struct {
	ID        int       `json:"ID"          form:"ID"          gorm:"column:id"`
	Nick      string    `json:"Nick"        form:"Nick"        gorm:"column:nick"`
	UserName  string    `json:"UserName"    form:"UserName"    gorm:"column:username"`
	Password  string    `json:"Password"    form:"Password"    gorm:"column:password"`
	Salt      string    `json:"Salt"        form:"Salt"        gorm:"column:salt"`
	Sex       int       `json:"Sex"         form:"Sex"         gorm:"column:sex"`
	HeadImage string    `json:"HeadImage"   form:"HeadImage"   gorm:"column:head_image"`
	Email     string    `json:"Email"       form:"Email"       gorm:"column:email"`
	Phone     string    `json:"Phone"       form:"Phone"       gorm:"column:phone"`
	Birthday  string    `json:"Birthday"    form:"Birthday"    gorm:"column:birthday"`
	Degree    string    `json:"Degree"      form:"Degree"      gorm:"column:degree"`
	Job       string    `json:"Job"         form:"Job"         gorm:"column:job"`
	Address   string    `json:"Address"     form:"Address"     gorm:"column:Address"`
	Status    int       `json:"Status"      form:"Status"      gorm:"column:status"`
	CreatedAt time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
	UpdatedAt time.Time `json:"UpdateTime"  form:"UpdateTime"  gorm:"column:lastupdate"`
}

func (User) TableName() string {
	return "user"
}

type Job struct {
	ID          int       `json:"ID"           form:"ID"           gorm:"column:id"`
	Name        string    `json:"Name"         form:"Name"         gorm:"column:name"`
	Pay         string    `json:"Pay"          form:"Pay"          gorm:"column:pay"`
	IcoUrl      string    `json:"IcoUrl"       form:"IcoUrl"       gorm:"column:ico_url"`
	Company     string    `json:"Company"      form:"Company"      gorm:"column:company"`
	Scale       string    `json:"Scale"        form:"Scale"        gorm:"column:scale"`
	Description string    `json:"Description"  form:"Description"  gorm:"column:description"`
	CreatedAt   time.Time `json:"CreateTime"   form:"CreateTime"   gorm:"column:createtime"`
	UpdatedAt   time.Time `json:"UpdateTime"   form:"UpdateTime"   gorm:"column:lastupdate"`
}

func (Job) TableName() string {
	return "job"
}

type JobTagMap struct {
	ID        int       `json:"ID"          form:"ID"          gorm:"column:id"`
	JobId     int       `json:"JobId"       form:"JobId"       gorm:"column:job_id"`
	Tag       string    `json:"Tag"         form:"Tag"         gorm:"column:tag"`
	CreatedAt time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
}

func (JobTagMap) TableName() string {
	return "job_tag_map"
}

type Resume struct {
	ID             int    `json:"ID"              form:"ID"              gorm:"column:id"`
	UserId         int    `json:"UserId"          form:"UserId"          gorm:"column:user_id"`
	State          string `json:"State"           form:"State"           gorm:"column:state"`
	City           string `json:"City"            form:"City"            gorm:"column:city"`
	Identity       string `json:"Identity"        form:"Identity"        gorm:"column:identity"`
	Intention      string `json:"Intention"       form:"Intention"       gorm:"column:intention"`
	Advantage      string `json:"Advantage"       form:"Advantage"       gorm:"column:advantage"`
	WorkExperience string `json:"WorkExperience"  form:"WorkExperience"  gorm:"column:work_experience"`

	CreatedAt time.Time `json:"CreateTime"   form:"CreateTime"   gorm:"column:createtime"`
	UpdatedAt time.Time `json:"UpdateTime"   form:"UpdateTime"   gorm:"column:lastupdate"`
}

func (Resume) TableName() string {
	return "resume"
}
