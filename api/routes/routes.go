package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/madruga665/twitter-api/api/controllers"
)

func Routes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()

	v1 := router.Group("/v1")
	{
		v1.GET("/tweets", tweetController.GetAll)
		v1.GET("/tweet/:id", tweetController.GetById)
		v1.POST("/tweet", tweetController.Create)
		v1.DELETE("/tweet/:id", tweetController.Delete)
		v1.PUT("/tweet/:id", tweetController.Update)
	}

	return v1
}
