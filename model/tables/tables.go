package tables

import "time"

type User struct {
	ID        int       `json:"ID"          gorm:"column:id"`
	Nick      string    `json:"Nick"        gorm:"column:nick"`
	UserName  string    `json:"UserName"    gorm:"column:username"`
	Password  string    `json:"Password"    gorm:"column:password"`
	Salt      string    `json:"Salt"        gorm:"column:salt"`
	Sex       int       `json:"Sex"         gorm:"column:sex"`
	HeadImage string    `json:"HeadImage"   gorm:"column:head_image"`
	Email     string    `json:"Email"       gorm:"column:email"`
	Phone     string    `json:"Phone"       gorm:"column:phone"`
	Birthday  string    `json:"Birthday"    gorm:"column:birthday"`
	Degree    string    `json:"Degree"      gorm:"column:degree"`
	Job       string    `json:"Job"         gorm:"column:job"`
	Address   string    `json:"Address"     gorm:"column:Address"`
	Status    int       `json:"Status"      gorm:"column:status"`
	CreatedAt time.Time `json:"CreateTime"  gorm:"column:createtime"`
	UpdatedAt time.Time `json:"UpdateTime"  gorm:"column:lastupdate"`
}

func (User) TableName() string {
	return "user"
}

type Job struct {
	ID          int       `json:"ID"           gorm:"column:id"`
	Name        string    `json:"Name"         gorm:"column:name"`
	Pay         string    `json:"Pay"          gorm:"column:pay"`
	IcoUrl      string    `json:"IcoUrl"       gorm:"column:ico_url"`
	Company     string    `json:"Company"      gorm:"column:company"`
	Scale       string    `json:"Scale"        gorm:"column:scale"`
	Description string    `json:"Description"  gorm:"column:description"`
	CreatedAt   time.Time `json:"CreateTime"   gorm:"column:createtime"`
	UpdatedAt   time.Time `json:"UpdateTime"   gorm:"column:lastupdate"`
}

func (Job) TableName() string {
	return "job"
}

type JobTagMap struct {
	ID        int       `json:"ID"          gorm:"column:id"`
	JobId     int       `json:"JobId"       gorm:"column:job_id"`
	Tag       string    `json:"Tag"         gorm:"column:tag"`
	CreatedAt time.Time `json:"CreateTime"  gorm:"column:createtime"`
}

func (JobTagMap) TableName() string {
	return "job_tag_map"
}

type Resume struct {
	ID             int       `json:"ID"              gorm:"column:id"`
	UserId         int       `json:"UserId"          gorm:"column:user_id"`
	State          string    `json:"State"           gorm:"column:state"`
	City           string    `json:"City"            gorm:"column:city"`
	Identity       string    `json:"Identity"        gorm:"column:identity"`
	Intention      string    `json:"Intention"       gorm:"column:intention"`
	Advantage      string    `json:"Advantage"       gorm:"column:advantage"`
	WorkExperience string    `json:"WorkExperience"  gorm:"column:work_experience"`
	CreatedAt      time.Time `json:"CreateTime"      gorm:"column:createtime"`
	UpdatedAt      time.Time `json:"UpdateTime"      gorm:"column:lastupdate"`
}

func (Resume) TableName() string {
	return "resume"
}

type UserEducationMap struct {
	ID                int       `json:"ID"                 gorm:"column:id"`
	UserId            int       `json:"UserId"             gorm:"column:user_id"`
	SchoolName        string    `json:"SchoolName"         gorm:"column:school_name"`
	Major             string    `json:"Major"              gorm:"column:major"`
	MatriculationTime string    `json:"MatriculationTime"  gorm:"column:matriculation_time"`
	GraduationTime    string    `json:"GraduationTime"     gorm:"column:graduation_time"`
	CreatedAt         time.Time `json:"CreateTime"         gorm:"column:createtime"`
	UpdatedAt         time.Time `json:"UpdateTime"         gorm:"column:lastupdate"`
}

func (UserEducationMap) TableName() string {
	return "user_education_map"
}

type JobExpectation struct {
	ID        int       `json:"ID"          gorm:"column:id"`
	UserId    int       `json:"UserId"      gorm:"column:user_id"`
	JobTags   string    `json:"JobTags"     gorm:"column:job_tags"`
	Pay       string    `json:"Pay"         gorm:"column:pay"`
	City      string    `json:"City"        gorm:"column:city"`
	CreatedAt time.Time `json:"CreateTime"  gorm:"column:createtime"`
	UpdatedAt time.Time `json:"UpdateTime"  gorm:"column:lastupdate"`
}

func (JobExpectation) TableName() string {
	return "job_expectation"
}

type DeliverRecord struct {
	ID        int       `json:"ID"          gorm:"column:id"`
	UserId    int       `json:"UserId"      gorm:"column:user_id"`
	JobId     int       `json:"JobId"       gorm:"column:job_id"`
	CreatedAt time.Time `json:"CreateTime"  gorm:"column:createtime"`
}

func (DeliverRecord) TableName() string {
	return "deliver_record"
}
