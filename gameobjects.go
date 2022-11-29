package main

import "github.com/google/uuid"

type GameState struct {
	Board    *Board `json:"board"`
	Lobby    *Lobby `json:"lobby"`
	Winner   string `json:"winner"`
	Playable bool   `json:"playable"`
}

type Board struct {
	NextToPlay string  `json:"nextToPlay"`
	Tiles      []*Tile `json:"tiles"`
}

type Tile struct {
	TileLocation int    `json:"location"`
	TileOwner    string `json:"owner"`
}

type Lobby struct {
	LobbyId string   `json:"lobbyId"`
	Players []string `json:"players"`
}

func CreateNewGame(playerId string) (GameState, string) {
	lobbyId := uuid.New().String()
	return GameState{
		Board: &Board{
			NextToPlay: playerId,
			Tiles:      CreateTiles(),
		},
		Playable: false,
		Winner:   "null",
		Lobby: &Lobby{
			LobbyId: lobbyId,
			Players: []string{
				playerId,
			},
		},
	}, lobbyId
}

func CreateTiles() []*Tile {
	tileArray := []*Tile{}
	for i := 0; i < 9; i++ {
		tileArray = append(tileArray, &Tile{
			TileLocation: i,
			TileOwner:    "null",
		})
	}
	return tileArray
}
