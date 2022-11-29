package main

type MakeMoveRequest struct {
}

type GameStateResponse struct {
}

type CreateLobbyRequest struct {
	PlayerId string `json:"playerId"`
}

type CreateLobbyResponse struct {
}

type JoinLobbyRequest struct {
}
