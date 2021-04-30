package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"job-hunting/service/controller"
)

func Init() *gin.Engine {

	router := gin.Default()
	router.Use(cors.Default())

	Controller := new(controller.Controller)
	router.POST("/job-hunting/api", Controller.Handle)

	return router
}
