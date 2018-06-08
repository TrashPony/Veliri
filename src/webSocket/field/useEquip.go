package field

import "github.com/gorilla/websocket"

func UseEquip(msg Message, ws *websocket.Conn)  {
	client, findClient := usersFieldWs[ws]
	_, findUnit := client.GetUnit(msg.X, msg.Y)
	_, findGame := Games[client.GetGameID()]
	playerEquip, findEquip := client.GetEquipByID(msg.EquipID)

	//TODO 1) активация эфектов от эквипа на юнит
	//TODO 2) эквим делаем заюзаным
	//TODO 3) обновляем бд
	//TODO 4) оповещаем юзера об успешности операции и обновляем инфу на клиенте
	//TODO 5) оповещаем других игроков которые видят этого юнита
	//TODO 6) на фронтенде проигрывается анимация

	if findClient && findUnit && findGame && !client.GetReady() && findEquip {
		println(playerEquip.Type)
	} else {
		ws.WriteJSON(ErrorMessage{Event:msg.Event, Error: "not allow"})
	}
}
