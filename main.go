package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type song struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Artist    string `json:"artist"`
	PlayedNum int    `json:"playedNum"`
	Length    int    `json:"length"` //in seconds
}

var songList = []song{
	{ID: "1", Name: "Move in week", Artist: "Will Horizony", PlayedNum: 10, Length: 76},
	{ID: "2", Name: "Move in week 2", Artist: "Will Horizony", PlayedNum: 6, Length: 99},
	{ID: "3", Name: "Move in week 3", Artist: "Will Horizony", PlayedNum: 1, Length: 100},
}

func getSongs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, songList)
}

func addSong(context *gin.Context) {
	var newSong song
	// Take the json and convert it to song type, otherwise error
	if err := context.BindJSON(&newSong); err != nil {
		return
	}

	songList = append(songList, newSong)

	context.IndentedJSON(http.StatusCreated, newSong)
}

func getSongByID(id string) (*song, error) {
	for i, t := range songList {
		if t.ID == id {
			return &songList[i], nil
		}
	}

	return nil, errors.New("Song not found.")
}

func getSong(context *gin.Context) {
	id := context.Param("id")
	song, err := getSongByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested Song Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, song)
}

func updatePlayedNum(context *gin.Context) {
	id := context.Param("id")
	song, err := getSongByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested Song Not Found"})
		return
	}

	song.PlayedNum++

	context.IndentedJSON(http.StatusOK, song)
}

func deleteSongByID(id string) (*song, error) {
	for i, t := range songList {
		if t.ID == id {
			var song = &songList[i]
			songList[i] = songList[len(songList)-1]
			songList = songList[:len(songList)-1]
			return song, nil
		}
	}

	return nil, errors.New("Song not found.")
}

func deleteSong(context *gin.Context) {
	id := context.Param("id")
	song, err := deleteSongByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested Song Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, song)
}

func main() {
	router := gin.Default() // This IS the server
	router.GET("/songs", getSongs)
	router.GET("/songs/:id", getSong)
	router.PATCH("/songs/:id", updatePlayedNum)
	router.DELETE("/songs/:id", deleteSong)
	router.POST("/songs", addSong)
	router.Run("localhost:9090")
}
