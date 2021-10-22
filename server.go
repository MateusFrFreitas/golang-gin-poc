package main

import (
	"net/http"

	"github.com/MateusFrFreitas/goland-gin-poc/controller"
	service "github.com/MateusFrFreitas/goland-gin-poc/services"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})

	server.Run(":8080")
}
