package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/useEquip"
	"../../mechanics/unit"
	"../../mechanics/equip"
	"../../mechanics/player"
	"../../mechanics/game"

)

func UseEquip(msg Message, ws *websocket.Conn) {
	client, findClient := usersFieldWs[ws]
	activeGame, findGame := Games[client.GetGameID()]
	playerEquip, findEquip := client.GetEquipByID(msg.EquipID)

	//TODO 1) активация эфектов от эквипа на юнит +
	//TODO 1,5 ) активация эфектов от эквипа на территорию
	//TODO 2) эквим делаем заюзаным +
	//TODO 3) обновляем бд +
	//TODO 4) оповещаем юзера об успешности операции и обновляем инфу на клиенте
	//TODO 5) оповещаем других игроков которые видят этого юнита
	//TODO 6) на фронтенде проигрывается анимация

	if findClient && findGame && !client.GetReady() && findEquip && !playerEquip.Used {
		if playerEquip.Applicable == "map" {
			gameCoordinate, findCoordinate := client.GetWatchCoordinate(msg.X, msg.Y)
			if findCoordinate {
				useEquip.ToMap(gameCoordinate, activeGame, playerEquip, client)
			} else {
				ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not find coordinate"})
			}
		} else {
			gameUnit, findUnit := client.GetUnit(msg.X, msg.Y)
			if findUnit {
				// todo свои, чужие юниты
				useEquip.ToUnit(gameUnit, playerEquip, client)
				ws.WriteJSON(SendUseEquip{Event: msg.Event, Unit: gameUnit, AppliedEquip: playerEquip})
				updateEquipHostileUser(client, activeGame, gameUnit, playerEquip)
			} else {
				ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not find unit"})
			}
		}
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "not allow"})
	}
}

func updateEquipHostileUser(client *player.Player, activeGame *game.Game, gameUnit *unit.Unit, playerEquip *equip.Equip) {
	for _, user := range activeGame.GetPlayers() {
		if user.GetLogin() != client.GetLogin() {
			_, watch := user.GetHostileUnit(gameUnit.X, gameUnit.Y)
			if watch {
				equipPipe <- SendUseEquip{Event: "UseEquip", UserName: user.GetLogin(), GameID: activeGame.Id, Unit: gameUnit, AppliedEquip: playerEquip}
			}
		}
	}
}

type SendUseEquip struct {
	Event        string       `json:"event"`
	UserName     string       `json:"user_name"`
	GameID       int          `json:"game_id"`
	Unit         *unit.Unit   `json:"unit"`
	AppliedEquip *equip.Equip `json:"applied_equip"`
}
