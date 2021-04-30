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
	Job       string    `json:"Job"         form:"Job"         gorm:"column:job"`
	Address   string    `json:"Address"     form:"Address"     gorm:"column:Address"`
	Status    int       `json:"Status"      form:"Status"      gorm:"column:status"`
	CreatedAt time.Time `json:"CreateTime"  form:"CreateTime"  gorm:"column:createtime"`
	UpdatedAt time.Time `json:"UpdateTime"  form:"UpdateTime"  gorm:"column:lastupdate"`
}

func (User) TableName() string {
	return "user"
}
