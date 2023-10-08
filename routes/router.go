package routes

import (
	"Assignment-3/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/sensor", controller.GetSensorData)
	return r
}
