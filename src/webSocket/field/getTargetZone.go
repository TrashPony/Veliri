package field

import "github.com/gorilla/websocket"

func GetTargetZone(msg Message, ws *websocket.Conn)  {

	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.X, msg.Y)
	activeGame, findGame := Games[client.GetGameID()]

	if findClient && findUnit && findGame {

		tmpUnit := *gameUnit

		tmpUnit.SetX(msg.ToX)
		tmpUnit.SetY(msg.ToY)

		SelectTarget(client, &tmpUnit, activeGame, ws)
	}
}
