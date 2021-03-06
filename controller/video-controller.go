package controller

import (
	"net/http"

	"github.com/MateusFrFreitas/goland-gin-poc/customvalidator"
	"github.com/MateusFrFreitas/goland-gin-poc/entity"
	"github.com/MateusFrFreitas/goland-gin-poc/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", customvalidator.ValidateCoolTitle)

	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video

	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	c.service.Save(video)

	return nil
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
