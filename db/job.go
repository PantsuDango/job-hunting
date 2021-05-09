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

func (SocialDB) SelectJob(Offset, Limit int, Keyword string) ([]params.JobParams, int) {
	var job []params.JobParams
	var count int
	if Limit == 0 {
		Limit = 10
	}
	if len(Keyword) > 0 {
		Keyword := "%" + Keyword + "%"
		exeDB.Where("name like ?", Keyword).Offset(Offset).Limit(Limit).Order("createtime desc").Find(&job)
	} else {
		exeDB.Offset(Offset).Limit(Limit).Order("createtime desc").Find(&job)
	}
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

func (SocialDB) SelectDeliverRecordCount(user_id int) int {
	var count int
	exeDB.Model(&tables.DeliverRecord{}).Where("user_id = ?", user_id).Count(&count)
	return count
}

func (SocialDB) SelectResume(user_id int) tables.Resume {
	var resume tables.Resume
	exeDB.Where("user_id = ?", user_id).Find(&resume)
	return resume
}

func (SocialDB) SelectUserEducationMap(user_id int) tables.UserEducationMap {
	var user_education_map tables.UserEducationMap
	exeDB.Where("user_id = ?", user_id).Find(&user_education_map)
	return user_education_map
}

func (SocialDB) SelectJobExpectation(user_id int) tables.JobExpectation {
	var job_expectation tables.JobExpectation
	exeDB.Where("user_id = ?", user_id).Find(&job_expectation)
	return job_expectation
}

func (SocialDB) UpdateUser(user tables.User) error {
	err := exeDB.Save(&user).Error
	return err
}

func (SocialDB) UpdateResume(resume tables.Resume) error {
	err := exeDB.Save(&resume).Error
	return err
}

func (SocialDB) UpdateUserEducationMap(user_education_map tables.UserEducationMap) error {
	err := exeDB.Save(&user_education_map).Error
	return err
}

func (SocialDB) UpdateJobExpectation(job_expectation tables.JobExpectation) error {
	err := exeDB.Save(&job_expectation).Error
	return err
}
