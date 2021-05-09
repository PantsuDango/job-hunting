package controller

import (
	"crypto/md5"
	"fmt"
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

// 修改简历详情
func (Controller Controller) ModifyResume(ctx *gin.Context, user tables.User) {

	// 校验前端传的参数是否符合预期
	var ResumeInfoParams result.ResumeInfo
	if err := ctx.ShouldBindBodyWith(&ResumeInfoParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	// 查询简历详情
	resume := Controller.SocialDB.SelectResume(user.ID)
	resume.UserId = user.ID
	// 改所在城市
	if len(ResumeInfoParams.City) > 0 {
		resume.City = ResumeInfoParams.City
	}
	// 改身份
	if len(ResumeInfoParams.Identity) > 0 {
		resume.Identity = ResumeInfoParams.Identity
	}
	// 改求职状态
	if len(ResumeInfoParams.State) > 0 {
		resume.State = ResumeInfoParams.State
	}
	// 改个人优势
	if len(ResumeInfoParams.Advantage) > 0 {
		resume.Advantage = ResumeInfoParams.Advantage
	}
	// 改求职意向
	if len(ResumeInfoParams.Intention) > 0 {
		resume.Intention = ResumeInfoParams.Intention
	}
	// 改工作经历
	if len(ResumeInfoParams.WorkExperience) > 0 {
		resume.WorkExperience = ResumeInfoParams.WorkExperience
	}

	// 修改简历详情
	err := Controller.SocialDB.UpdateResume(resume)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "Access resume table error.", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}

// 修改教育经历
func (Controller Controller) ModifyEducation(ctx *gin.Context, user tables.User) {

	// 校验前端传的参数是否符合预期
	var EducationInfoParams result.EducationInfo
	if err := ctx.ShouldBindBodyWith(&EducationInfoParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	// 查询教育经历
	user_education_map := Controller.SocialDB.SelectUserEducationMap(user.ID)
	user_education_map.UserId = user.ID
	// 改毕业时间
	if len(EducationInfoParams.GraduationTime) > 0 {
		user_education_map.GraduationTime = EducationInfoParams.GraduationTime
	}
	// 改专业名称
	if len(EducationInfoParams.Major) > 0 {
		user_education_map.Major = EducationInfoParams.Major
	}
	// 改入学时间
	if len(EducationInfoParams.MatriculationTime) > 0 {
		user_education_map.MatriculationTime = EducationInfoParams.MatriculationTime
	}
	// 改学校名称
	if len(EducationInfoParams.SchoolName) > 0 {
		user_education_map.SchoolName = EducationInfoParams.SchoolName
	}

	// 修改教育经历
	err := Controller.SocialDB.UpdateUserEducationMap(user_education_map)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "Access user_education_map table error.", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}

// 修改求职期望
func (Controller Controller) ModifyJobExpectation(ctx *gin.Context, user tables.User) {

	// 校验前端传的参数是否符合预期
	var JobExpectationParams result.JobExpectation
	if err := ctx.ShouldBindBodyWith(&JobExpectationParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	// 查询求职期望
	job_expectation := Controller.SocialDB.SelectJobExpectation(user.ID)
	job_expectation.UserId = user.ID
	// 改期望城市
	if len(JobExpectationParams.City) > 0 {
		job_expectation.City = JobExpectationParams.City
	}
	// 改职位类别
	if len(JobExpectationParams.JobTags) > 0 {
		job_expectation.JobTags = JobExpectationParams.JobTags
	}
	// 改期望薪资
	if len(JobExpectationParams.Pay) > 0 {
		job_expectation.Pay = JobExpectationParams.Pay
	}
	// 修改求职期望
	err := Controller.SocialDB.UpdateJobExpectation(job_expectation)
	if err != nil {
		JSONFail(ctx, http.StatusOK, AccessDBError, "Access job_expectation table error.", gin.H{
			"Code":    AccessDBError,
			"Message": err.Error(),
		})
		return
	}

	JSONSuccess(ctx, http.StatusOK, "Success")
}

// 修改密码
func (Controller Controller) ModifyPassword(ctx *gin.Context, user tables.User) {

	// 校验前端传的参数是否符合预期
	var ModifyPasswordParams params.ModifyPasswordParams
	if err := ctx.ShouldBindBodyWith(&ModifyPasswordParams, binding.JSON); err != nil {
		JSONFail(ctx, http.StatusOK, IllegalRequestParameter, "Invalid json or illegal request parameter", gin.H{
			"Code":    IncompleteParameters,
			"Message": err.Error(),
		})
		return
	}

	// 检验旧密码
	s := ModifyPasswordParams.OldPassword + user.Salt
	if fmt.Sprintf("%x", md5.Sum([]byte(s))) == user.Password {
		user.Salt = GetRandomString(8)
		s := ModifyPasswordParams.NewPassword + user.Salt
		user.Password = fmt.Sprintf("%x", md5.Sum([]byte(s)))
	} else {
		JSONFail(ctx, http.StatusOK, PasswordError, "OldPassword error.", gin.H{
			"Code":    "InvalidJSON",
			"Message": "OldPassword error.",
		})
		return
	}

	// 修改密码
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
