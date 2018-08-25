package field

import "github.com/gorilla/websocket"

func SelectEquip(msg Message, ws *websocket.Conn)  {
	println(msg.X , msg.Y)
	println(msg.EquipType , msg.NumberSlot)

	/*client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.X, msg.Y)
	activeGame, findGame := Games.Get(client.GetGameID())

	if findClient && findUnit && findGame {

	}*/
}
