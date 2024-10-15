package controllers

import (
	"fmt"
	entities "go-docker-postgres/api/entities"
	database "go-docker-postgres/db"

	"net/http"

	"github.com/gin-gonic/gin"
)

var db = database.ConnectToDatabase()

func FindAll(ctx *gin.Context) {
	rows, err := db.Query("SELECT * FROM tweets")

	if err != nil {
		fmt.Println("Erro ao executar a query:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tweets []entities.Tweet
	for rows.Next() {
		var tweet entities.Tweet
		err := rows.Scan(&tweet.ID, &tweet.Content)
		if err != nil {
			fmt.Println("Erro ao escanear os resultados:", err)
			continue
		}
		tweets = append(tweets, tweet)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Erro geral ao iterar sobre as linhas:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tweets)
}

func Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	_, err := db.Exec("INSERT INTO tweets (id, content) values ($1, $2)", tweet.ID, tweet.Content)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := db.Exec("DELETE FROM tweets where id = $1", id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
