package lobby

import (
	"../../mechanics/player"
)

func RefreshLobbyGames(user *player.Player) {
	for _, client := range usersLobbyWs {
		if client.GetLogin() != user.GetLogin() {

			var refresh = Response{Event: "GameRefresh", UserName: client.GetLogin()}
			lobbyPipe <- refresh

			for _, game := range openGames {

				var resp = Response{Event: "GameView", UserName: client.GetLogin(), Game: game}
				lobbyPipe <- resp

			}
		}
	}
}
