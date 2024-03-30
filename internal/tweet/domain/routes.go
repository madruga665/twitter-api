package routes

import (
	"github.com/gin-gonic/gin"
	connection "github.com/madruga665/twitter-api/db"
	"github.com/madruga665/twitter-api/internal/tweet/domain/controller"
	"github.com/madruga665/twitter-api/internal/tweet/domain/repository"
	"github.com/madruga665/twitter-api/internal/tweet/domain/service"
)

func Routes(router *gin.Engine) *gin.RouterGroup {
	db := connection.DB
	repository := repository.New(db)
	service := service.New(repository)
	controlller := controller.New(service)

	v1 := router.Group("/v1")
	{
		v1.GET("/tweets", controlller.GetAll)
		v1.GET("/tweet/:id", controlller.GetById)
		v1.POST("/tweet", controlller.Create)
		v1.DELETE("/tweet/:id", controlller.Delete)
		v1.PUT("/tweet/:id", controlller.Update)
	}

	return v1
}
