package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"job-hunting/model/params"
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
