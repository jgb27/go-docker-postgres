package entities

import "github.com/google/uuid"

type Tweet struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func NewTweet() *Tweet {
	Tweet := Tweet{
		ID: uuid.New().String(),
	}

	return &Tweet
}
