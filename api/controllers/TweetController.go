package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madruga665/twitter-api/api/entities"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (controller *tweetController) GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, controller.tweets)
}

func (controller *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if error := ctx.BindJSON(&tweet); error != nil {
		log.Fatal("Erro ao criar tweet")
		return
	}

	controller.tweets = append(controller.tweets, *tweet)

	ctx.JSON(http.StatusNoContent, controller.tweets)
}

func (controller *tweetController) Delete(ctx *gin.Context) {
	tweetId := ctx.Param("id")

	for index, tweet := range controller.tweets {
		if tweet.ID == tweetId {
			controller.tweets = append(controller.tweets[0:index], controller.tweets[index+1:]...)
		}
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Tweet not found",
	})
}

func (controller *tweetController) GetById(ctx *gin.Context) {
	tweetId := ctx.Param("id")

	for _, tweet := range controller.tweets {
		if tweet.ID == tweetId {
			ctx.JSON(http.StatusOK, tweet)
		}
	}
}

func (controller *tweetController) Update(ctx *gin.Context) {
	tweetId := ctx.Param("id")
	body, _ := io.ReadAll(ctx.Request.Body)

	var jsonData map[string]interface{}

	json.Unmarshal(body, &jsonData)

	descriptionBody := jsonData["description"].(string)

	for _, tweet := range controller.tweets {
		if tweet.ID == tweetId {
			tweet.Description = descriptionBody
			ctx.JSON(http.StatusOK, tweet)
			break
		}
	}
}
