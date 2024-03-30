package repository

import "database/sql"

type TweetRepositoryInterface interface {
	GetAll() (*sql.Rows, error)
	GetById(tweetId string) (*sql.Rows, error)
}
