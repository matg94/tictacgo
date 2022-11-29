package main

type MakeMoveRequest struct {
	TileLocation int `json:"tileLocation"`
}

type CreateLobbyRequest struct {
	PlayerId string `json:"playerId"`
}

type JoinLobbyRequest struct {
	PlayerId string `json:"playerId"`
}
