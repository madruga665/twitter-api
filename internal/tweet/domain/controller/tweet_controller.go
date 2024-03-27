package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	connection "github.com/madruga665/twitter-api/db"
	"github.com/madruga665/twitter-api/internal/tweet/domain/entity"
	"github.com/madruga665/twitter-api/internal/tweet/domain/repository"
)

type tweetController struct{}

func NewTweetController() tweetController {
	return tweetController{}
}

func (controller tweetController) GetAll(ctx *gin.Context) {
	db := connection.DB
	tweets, _ := repository.GetAll(db)

	ctx.JSON(http.StatusOK, tweets)
}

func (controller tweetController) GetById(ctx *gin.Context) {
	db := connection.DB
	tweetId := ctx.Param("id")
	tweet, _ := repository.GetById(db, tweetId)

	ctx.JSON(http.StatusOK, tweet)
}

func (controller tweetController) Create(ctx *gin.Context) {
	db := connection.DB
	tweet := entity.NewTweet()

	if error := ctx.BindJSON(&tweet); error != nil {
		log.Fatal("Erro ao criar tweet")
		return
	}

	repository.Create(db, *tweet)

	ctx.JSON(http.StatusCreated, tweet)
}

func (controller tweetController) Update(ctx *gin.Context) {
	db := connection.DB
	tweetId := ctx.Param("id")

	var tweet entity.UpdateTweet
	if err := ctx.ShouldBindJSON(&tweet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repository.Update(db, tweetId, tweet.Description)

	ctx.JSON(http.StatusNoContent, nil)
}

func (controller tweetController) Delete(ctx *gin.Context) {
	db := connection.DB
	tweetId := ctx.Param("id")
	repository.Delete(db, tweetId)

	ctx.JSON(http.StatusNoContent, nil)
}
