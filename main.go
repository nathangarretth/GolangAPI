package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type song struct {
	ID        string `json:"Id"`
	Name      string `json:"Name"`
	Artist    string `json:"Artist"`
	PlayedNum int    `json:"PlayedNum"`
	Length    int    `json:"Length"`
}

var songList = []song{
	{ID: "1", Name: "Move in week", Artist: "Will Horizony", PlayedNum: 10, Length: 76},
	{ID: "2", Name: "Move in week 2", Artist: "Will Horizony", PlayedNum: 10, Length: 76},
	{ID: "3", Name: "Move in week 3", Artist: "Will Horizony", PlayedNum: 10, Length: 76},
}

func getSongs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, songList)
}

func main() {
	router := gin.Default()
	router.GET("/songs")
	router.Run("localhost:9090")
}
