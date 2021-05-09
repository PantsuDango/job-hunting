package db

import (
	"job-hunting/model/params"
	"job-hunting/model/tables"
)

type SocialDB struct{}

func (SocialDB) GetUserInfo(UserName string) (tables.User, error) {
	var user tables.User
	err := exeDB.Where("username = ? AND status= 0 ", UserName).First(&user).Error
	return user, err
}

func (SocialDB) QueryUserById(userId int) (tables.User, error) {
	var user tables.User
	err := exeDB.Where("id = ? AND status = 0", userId).First(&user).Error
	return user, err
}

func (SocialDB) CreateUser(user tables.User) error {
	err := exeDB.Create(&user).Error
	return err
}

func (SocialDB) CreateJob(job *tables.Job) error {
	err := exeDB.Create(&job).Error
	return err
}

func (SocialDB) SelectJob(Offset, Limit int) ([]params.JobParams, int) {
	var job []params.JobParams
	var count int
	if Limit == 0 {
		Limit = 10
	}
	exeDB.Offset(Offset).Limit(Limit).Order("createtime desc").Find(&job)
	exeDB.Model(&[]params.JobParams{}).Count(&count)
	return job, count
}

func (SocialDB) CreateJobTagMap(job_tag_map tables.JobTagMap) error {
	err := exeDB.Create(&job_tag_map).Error
	return err
}

func (SocialDB) SelectJobTagMapByJobId(job_id int) []tables.JobTagMap {
	var job_tag_map []tables.JobTagMap
	exeDB.Where("job_id = ?", job_id).Find(&job_tag_map)
	return job_tag_map
}

func (SocialDB) SelectJobById(id int) (params.JobParams, error) {
	var job params.JobParams
	err := exeDB.Where("id = ?", id).Find(&job).Error
	return job, err
}

func (SocialDB) SelectDeliverRecord(job_id, user_id int) int {
	var count int
	exeDB.Model(&tables.DeliverRecord{}).Where("job_id = ? AND user_id = ?", job_id, user_id).Count(&count)
	return count
}

func (SocialDB) CreateDeliverRecord(deliver_record tables.DeliverRecord) error {
	err := exeDB.Create(&deliver_record).Error
	return err
}
