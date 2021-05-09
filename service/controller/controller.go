package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"job-hunting/model/params"
	"job-hunting/model/result"
	"job-hunting/model/tables"
	"net/http"
)

// 新建职位
func (Controller Controller) AddJob(ctx *gin.Context, user tables.User) {

	// 校验前端传的参数是否符合预期
	var JobParams params.JobParams
	if err := ctx.ShouldBindBodyWith(&JobParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	// 新建职位记录
	var Job tables.Job
	Job.Name = JobParams.Name
	Job.Description = JobParams.Description
	Job.Pay = JobParams.Pay
	Job.IcoUrl = JobParams.IcoUrl
	Job.Company = JobParams.Company
	Job.Scale = JobParams.Scale
	err := Controller.SocialDB.CreateJob(&Job)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "Access job table error.", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	// 新建岗位与标签映射记录
	for _, tag := range JobParams.Tags {
		var job_tag_map tables.JobTagMap
		job_tag_map.JobId = Job.ID
		job_tag_map.Tag = tag
		err = Controller.SocialDB.CreateJobTagMap(job_tag_map)
	}
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "Access job_tag_map table error.", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}

// 查询岗位列表
func (Controller Controller) ListJob(ctx *gin.Context, user tables.User) {

	// 校验前端传的参数是否符合预期
	var ListJobParams params.ListJobParams
	if err := ctx.ShouldBindBodyWith(&ListJobParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	// 查询工作岗位
	jobs, count := Controller.SocialDB.SelectJob(ListJobParams.Offset, ListJobParams.Limit, ListJobParams.Keyword)
	for index, _ := range jobs {
		jobs[index].Createtime = jobs[index].CreatedAt.Format("2006-01-02 15:04:05")
		// 查询岗位标签
		job_tag_map := Controller.SocialDB.SelectJobTagMapByJobId(jobs[index].ID)
		for _, tmp := range job_tag_map {
			jobs[index].Tags = append(jobs[index].Tags, tmp.Tag)
		}
		// 查询是否投递过当前职位
		count := Controller.SocialDB.SelectDeliverRecord(jobs[index].ID, user.ID)
		if count > 0 {
			jobs[index].Isdeliver = true
		}
	}

	var ListJob result.ListJob
	ListJob.ListJob = jobs
	ListJob.Count = count

	JSONSuccess(ctx, http.StatusOK, ListJob)
}

// 查询某岗位详情
func (Controller Controller) JobInfo(ctx *gin.Context, user tables.User) {

	// 校验前端传的参数是否符合预期
	var JobInfoParams params.JobInfoParams
	if err := ctx.ShouldBindBodyWith(&JobInfoParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	// 查询某岗位详情
	job, err := Controller.SocialDB.SelectJobById(JobInfoParams.ID)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "Access job table error.", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}
	job.Createtime = job.CreatedAt.Format("2006-01-02 15:04:05")
	job_tag_map := Controller.SocialDB.SelectJobTagMapByJobId(job.ID)
	for _, tmp := range job_tag_map {
		job.Tags = append(job.Tags, tmp.Tag)
	}
	// 查询是否投递过当前职位
	count := Controller.SocialDB.SelectDeliverRecord(job.ID, user.ID)
	if count > 0 {
		job.Isdeliver = true
	}

	JSONSuccess(ctx, http.StatusOK, job)
}

// 简历投递
func (Controller Controller) DeliverJob(ctx *gin.Context, user tables.User) {

	// 校验前端传的参数是否符合预期
	var JobInfoParams params.JobInfoParams
	if err := ctx.ShouldBindBodyWith(&JobInfoParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	// 检验岗位是否存在
	_, err := Controller.SocialDB.SelectJobById(JobInfoParams.ID)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "Access job table error.", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	// 创建岗位投递记录
	var deliver_record tables.DeliverRecord
	deliver_record.UserId = user.ID
	deliver_record.JobId = JobInfoParams.ID
	err = Controller.SocialDB.CreateDeliverRecord(deliver_record)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "Access deliver_record table error.", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}

// 用户个人信息
func (Controller Controller) UserInfo(ctx *gin.Context, user tables.User) {

	var UserInfo result.UserInfo

	// 用户个人信息
	UserInfo.UserInfo.ID = user.ID
	UserInfo.UserInfo.Sex = user.Sex
	UserInfo.UserInfo.Job = user.Job
	UserInfo.UserInfo.HeadImage = user.HeadImage
	UserInfo.UserInfo.Phone = user.Phone
	UserInfo.UserInfo.Email = user.Email
	UserInfo.UserInfo.Nick = user.Nick
	UserInfo.UserInfo.Address = user.Address
	UserInfo.UserInfo.Birthday = user.Birthday
	UserInfo.UserInfo.Degree = user.Degree
	UserInfo.UserInfo.UserName = user.UserName
	// 用户投递简历数
	UserInfo.DeliverCount = Controller.SocialDB.SelectDeliverRecordCount(user.ID)
	// 查询简历详情
	resume := Controller.SocialDB.SelectResume(user.ID)
	UserInfo.ResumeInfo.City = resume.City
	UserInfo.ResumeInfo.Identity = resume.Identity
	UserInfo.ResumeInfo.State = resume.State
	UserInfo.ResumeInfo.Advantage = resume.Advantage
	UserInfo.ResumeInfo.Intention = resume.Intention
	UserInfo.ResumeInfo.WorkExperience = resume.WorkExperience
	// 查询教育经历
	user_education_map := Controller.SocialDB.SelectUserEducationMap(user.ID)
	UserInfo.EducationInfo.GraduationTime = user_education_map.GraduationTime
	UserInfo.EducationInfo.Major = user_education_map.Major
	UserInfo.EducationInfo.MatriculationTime = user_education_map.MatriculationTime
	UserInfo.EducationInfo.SchoolName = user_education_map.SchoolName
	// 查询求职期望
	job_expectation := Controller.SocialDB.SelectJobExpectation(user.ID)
	UserInfo.JobExpectation.City = job_expectation.City
	UserInfo.JobExpectation.JobTags = job_expectation.JobTags
	UserInfo.JobExpectation.Pay = job_expectation.Pay

	JSONSuccess(ctx, http.StatusOK, UserInfo)
}

// 修改用户个人信息
func (Controller Controller) ModifyUser(ctx *gin.Context, user tables.User) {

	// 校验前端传的参数是否符合预期
	var UserParams result.User
	if err := ctx.ShouldBindBodyWith(&UserParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}
	// 改性别
	if UserParams.Sex != user.Sex {
		user.Sex = UserParams.Sex
	}
	// 改职位
	if len(UserParams.Job) > 0 {
		user.Job = UserParams.Job
	}
	// 改头像
	if len(UserParams.HeadImage) > 0 {
		user.HeadImage = UserParams.HeadImage
	}
	// 改电话
	if len(UserParams.Phone) > 0 {
		user.Phone = UserParams.Phone
	}
	// 改邮箱
	if len(UserParams.Email) > 0 {
		user.Email = UserParams.Email
	}
	// 改姓名
	if len(UserParams.Nick) > 0 {
		user.Nick = UserParams.Nick
	}
	// 改家庭地址
	if len(UserParams.Address) > 0 {
		user.Address = UserParams.Address
	}
	// 改生日
	if len(UserParams.Birthday) > 0 {
		user.Birthday = UserParams.Birthday
	}
	// 改学历
	if len(UserParams.Degree) > 0 {
		user.Degree = UserParams.Degree
	}
	// 修改用户个人信息
	err := Controller.SocialDB.UpdateUser(user)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "Access user table error.", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}
