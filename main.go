package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var state *State

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello, world!",
	})
}

func CreateLobby(c *gin.Context) {
	body := c.Request.Body
	requestBody, readError := ioutil.ReadAll(body)
	if readError != nil {
		c.JSON(400, gin.H{"error": readError.Error()})
	}
	req := &CreateLobbyRequest{}
	err := json.Unmarshal(requestBody, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	game, lobbyId := CreateNewGame(req.PlayerId)
	state.Games = append(state.Games, game)
	c.JSON(
		200,
		gin.H{
			"lobbyId": lobbyId,
		},
	)
}

// func GetGameState(c *gin.Context) {
// 	lobbyId := c.Param("lobbyId")

// }

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	state = &State{
		[]GameState{},
	}
	router.GET("/hello", Hello)
	router.POST("/createLobby", CreateLobby)
	// router.GET("/gamestate/:lobbyId", GetGameState)
	err := router.Run()
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}

}
