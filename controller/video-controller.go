package controller

import (
	"github.com/MateusFrFreitas/goland-gin-poc/entity"
	"github.com/MateusFrFreitas/goland-gin-poc/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)

	c.service.Save(video)

	return video
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
