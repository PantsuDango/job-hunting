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
	jobs, count := Controller.SocialDB.SelectJob(ListJobParams.Offset, ListJobParams.Limit)
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
