package service

import "github.com/madruga665/twitter-api/internal/tweet/domain/entity"

type TweetServiceInterface interface {
	GetAll() []entity.Tweet
}
