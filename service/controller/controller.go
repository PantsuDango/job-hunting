package controller

import (
	"github.com/gin-gonic/gin"
	"job-hunting/model/tables"
	"net/http"
)

// 新建帖子
func (Controller Controller) Test(ctx *gin.Context, user tables.User) {
	JSONSuccess(ctx, http.StatusOK, user)
}
