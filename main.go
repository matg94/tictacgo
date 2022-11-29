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
	state.Games = append(state.Games, &game)
	c.JSON(
		200,
		gin.H{
			"lobbyId": lobbyId,
		},
	)
}

func GetGameState(c *gin.Context) {
	lobbyId := c.Param("lobbyId")
	gameState := state.FindGame(lobbyId)
	emptyState := GameState{}
	if *gameState == emptyState {
		c.JSON(400, gin.H{
			"error": "lobby not found!",
		})
	}
	c.JSON(200, gameState)
}

func JoinLobby(c *gin.Context) {
	body := c.Request.Body
	requestBody, readError := ioutil.ReadAll(body)
	if readError != nil {
		c.JSON(400, gin.H{"error": readError.Error()})
		return
	}
	req := &JoinLobbyRequest{}
	err := json.Unmarshal(requestBody, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	lobbyId := c.Param("lobbyId")
	gameState := state.FindGame(lobbyId)
	emptyState := GameState{}
	if *gameState == emptyState {
		c.JSON(400, gin.H{
			"error": "lobby not found!",
		})
		return
	}
	gameState.Lobby.Players = append(gameState.Lobby.Players, req.PlayerId)
	gameState.Playable = true
	c.JSON(200, gameState)
}

func MakeMove(c *gin.Context) {
	body := c.Request.Body
	requestBody, readError := ioutil.ReadAll(body)
	if readError != nil {
		c.JSON(400, gin.H{"error": readError.Error()})
		return
	}
	req := &MakeMoveRequest{}
	err := json.Unmarshal(requestBody, req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	lobbyId := c.Param("lobbyId")
	gameState := state.FindGame(lobbyId)
	emptyState := GameState{}
	if *gameState == emptyState {
		c.JSON(400, gin.H{
			"error": "lobby not found!",
		})
		return
	}
	if req.PlayerId != gameState.Board.NextToPlay {
		c.JSON(400, gin.H{
			"error": "Not your move!",
		})
		return
	}
	if !gameState.Playable {
		c.JSON(400, gin.H{
			"error": "Game not playable!",
		})
		return
	}
	res := MakeStateMove(gameState, req.TileLocation, req.PlayerId)
	if !res {
		c.JSON(400, gin.H{
			"error": "invalid move",
		})
		return
	}

	winner := CheckWinner(gameState)
	gameState.Winner = winner
	if winner != "null" {
		gameState.Playable = false
	}

	gameState.Board.NextToPlay = FindOtherPlayerName(gameState, req.PlayerId)

	c.JSON(200, gameState)
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	state = &State{
		[]*GameState{},
	}
	router.GET("/hello", Hello)
	router.POST("/createLobby", CreateLobby)
	router.POST("/joinLobby/:lobbyId", JoinLobby)
	router.POST("/makeMove/:lobbyId", MakeMove)
	router.GET("/gamestate/:lobbyId", GetGameState)
	err := router.Run()
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}

}
