package service

import (
	"log"

	"github.com/madruga665/twitter-api/internal/tweet/domain/entity"
	"github.com/madruga665/twitter-api/internal/tweet/domain/repository"
)

type tweetService struct {
	repository repository.TweetRepositoryInterface
}

func New(repository repository.TweetRepositoryInterface) *tweetService {
	return &tweetService{
		repository: repository,
	}
}

func (service *tweetService) GetAll() []entity.Tweet {
	result, error := service.repository.GetAll()

	if error != nil {
		log.Fatal("error ao recuperar tweets", error)
	}

	var tweets []entity.Tweet
	for result.Next() {
		var tweet entity.Tweet
		result.Scan(&tweet.ID, &tweet.Description)
		tweets = append(tweets, tweet)
	}

	return tweets
}
