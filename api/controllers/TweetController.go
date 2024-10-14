package controllers

import (
	entities "go-docker-postgres/api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (tc *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	tc.tweets = append(tc.tweets, *tweet)
	ctx.JSON(http.StatusOK, tc.tweets)
}

func (tc *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tc.tweets)
}

func (tc *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	found := false

	for index, tweet := range tc.tweets {
		if tweet.ID == id {
			tc.tweets = append(tc.tweets[:index], tc.tweets[index+1:]...)
			found = true
			break
		}
	}

	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Tweet not found"})
		return
	}

	ctx.JSON(http.StatusOK, tc.tweets)
}
