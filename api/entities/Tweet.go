package entities

import "github.com/google/uuid"

type Tweet struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type UpdateTweet struct {
	Description string `json:"description"`
}

func NewTweet() *Tweet {
	tweet := Tweet{
		ID: uuid.New().String(),
	}

	return &tweet
}
