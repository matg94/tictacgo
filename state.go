package main

type State struct {
	Games []GameState
}

func (state *State) FindGame(lobbyId string) GameState {
	for _, game := range state.Games {
		if game.Lobby.LobbyId == lobbyId {
			return game
		}
	}
	return GameState{}
}
