package main

func MakeStateMove(state *GameState, move int, playerId string) bool {
	for _, tile := range state.Board.Tiles {
		if tile.TileLocation == move {
			if tile.TileOwner != "null" {
				return false
			}
			tile.TileOwner = playerId
			return true
		}
	}
	return false
}

func FindOtherPlayerName(state *GameState, name string) string {
	for _, player := range state.Lobby.Players {
		if player != name {
			return player
		}
	}
	return name
}

func CheckWinner(state *GameState) string {
	if CheckDiaWin(state) != "null" {
		return CheckDiaWin(state)
	}
	if CheckHorizWin(state) != "null" {
		return CheckHorizWin(state)
	}
	if CheckVertWin(state) != "null" {
		return CheckVertWin(state)
	}
	return "null"
}

func CheckHorizWin(state *GameState) string {
	p1 := state.Lobby.Players[0]
	p2 := state.Lobby.Players[1]
	if CountByName(state, p1, 0, 1, 2) == 3 {
		return p1
	}
	if CountByName(state, p2, 0, 1, 2) == 3 {
		return p2
	}
	if CountByName(state, p1, 3, 4, 5) == 3 {
		return p1
	}
	if CountByName(state, p2, 3, 4, 5) == 3 {
		return p2
	}
	if CountByName(state, p1, 6, 7, 8) == 3 {
		return p1
	}
	if CountByName(state, p2, 6, 7, 8) == 3 {
		return p2
	}
	return "null"
}

func CheckVertWin(state *GameState) string {
	p1 := state.Lobby.Players[0]
	p2 := state.Lobby.Players[1]
	if CountByName(state, p1, 0, 3, 6) == 3 {
		return p1
	}
	if CountByName(state, p2, 0, 3, 6) == 3 {
		return p2
	}
	if CountByName(state, p1, 1, 4, 7) == 3 {
		return p1
	}
	if CountByName(state, p2, 1, 4, 7) == 3 {
		return p2
	}
	if CountByName(state, p1, 2, 5, 8) == 3 {
		return p1
	}
	if CountByName(state, p2, 2, 5, 8) == 3 {
		return p2
	}
	return "null"
}

func CheckDiaWin(state *GameState) string {
	p1 := state.Lobby.Players[0]
	p2 := state.Lobby.Players[1]
	if CountByName(state, p1, 0, 4, 8) == 3 {
		return p1
	}
	if CountByName(state, p2, 0, 4, 8) == 3 {
		return p2
	}
	if CountByName(state, p1, 6, 4, 2) == 3 {
		return p1
	}
	if CountByName(state, p2, 6, 4, 2) == 3 {
		return p2
	}
	return "null"
}

func CountByName(state *GameState, playerId string, tiles ...int) int {
	count := 0
	for _, tile := range state.Board.Tiles {
		if Contains(tiles, tile.TileLocation) && playerId == tile.TileOwner {
			count += 1
		}
	}
	return count
}

func Contains(arr []int, val int) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}
