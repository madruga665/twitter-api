package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madruga665/twitter-api/api/controllers/repositories"
	"github.com/madruga665/twitter-api/api/entities"
	connection "github.com/madruga665/twitter-api/db"
)

type tweetController struct{}

func NewTweetController() tweetController {
	return tweetController{}
}

func (controller tweetController) GetAll(ctx *gin.Context) {
	db := connection.DB
	tweets, _ := repositories.GetAll(db)

	ctx.JSON(http.StatusOK, tweets)
}

func (controller tweetController) GetById(ctx *gin.Context) {
	db := connection.DB
	tweetId := ctx.Param("id")
	tweet, _ := repositories.GetById(db, tweetId)

	ctx.JSON(http.StatusOK, tweet)
}

func (controller tweetController) Create(ctx *gin.Context) {
	db := connection.DB
	tweet := entities.NewTweet()

	if error := ctx.BindJSON(&tweet); error != nil {
		log.Fatal("Erro ao criar tweet")
		return
	}

	repositories.Create(db, *tweet)

	ctx.JSON(http.StatusCreated, nil)
}

func (controller tweetController) Update(ctx *gin.Context) {
	db := connection.DB
	tweetId := ctx.Param("id")

	var tweet entities.UpdateTweet
	if err := ctx.ShouldBindJSON(&tweet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repositories.Update(db, tweetId, tweet.Description)

	ctx.JSON(http.StatusNoContent, nil)
}

func (controller tweetController) Delete(ctx *gin.Context) {
	db := connection.DB
	tweetId := ctx.Param("id")
	repositories.Delete(db, tweetId)

	ctx.JSON(http.StatusNoContent, nil)
}
