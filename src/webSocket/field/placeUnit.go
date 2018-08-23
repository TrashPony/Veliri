package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/localGame/Phases/placePhase"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/localGame"
	"../../mechanics/localGame/map/watchZone"
	"strconv"
)

func placeUnit(msg Message, ws *websocket.Conn) {
	client, ok := usersFieldWs[ws]     // берем клиента/игрока которй нам кинул сообщение из карты подключений, ключем является ид подключения
	actionGame, ok := Games.Get(client.GetGameID()) // находим игру в которую он сейчас играет

	if client.GetReady() == false { // если клиент еще не нажал кнопку готов то идем дальше иначе отправляем сообщение о ошибке

		if !ok { // если не удалось взять подключение или игру то удалям игрока, тут неважая ошибка сам ищи кек
			delete(usersFieldWs, ws)
			return
		}

		storageUnit, find := client.GetUnitStorage(msg.UnitID) // ищем юнита у игрока которые находяться сейчас у него в трюме по id.
		// msg - это обьект сообщения которое к нам пришло .UnitID - это поле сообщения в нем нам с клиента пришел id юнита которое он хочет ставить. Смотри фаил fieldMessage.go
		if find { // если мы его нашли то идем дальше, если нет то кидаем ошибку о том что нас хотят наебать
			_, find = placePhase.GetPlaceCoordinate(client.GetSquad().MatherShip.X, client.GetSquad().MatherShip.Y,
				client.GetSquad().MatherShip.Body.RangeView, actionGame)[strconv.Itoa(msg.X)][strconv.Itoa(msg.Y)] // тут мы берем зону где можно строить
			// msg - это обьект сообщения которое к нам пришло. .Y и .X - это поле сообщения в нем нам с клиента пришли координаты куда он ставит юнита. Смотри фаил fieldMessage.go
			if find { // если координата куда хочет ставить юзер юнита есть в зоне строителства то идем дальше иначе кидаем ошибку о том что тут нельзя ставить
				_, find := actionGame.GetUnit(msg.X, msg.Y) // тут мы пытаемся взять юнита в игре на точке куда он хочет ставить, если юнит есть то туда ставить нельзя
				coordinate, _ := actionGame.Map.GetCoordinate(msg.X, msg.Y) // тут мы берем СУЩЕСТВУЮ координату на игровой карте

				if !find && coordinate.Type != "obstacle" { // если на точке нет юнита и там можно стоять то идем дальше
					err := placePhase.PlaceUnit(storageUnit, msg.X, msg.Y, actionGame, client) // todo тут мы идем в механику выбери PlaceUnit и нажми ctrl+B <-- тебе сюда
					// todo тут ты должен получить обьект пути и отправить его на фронтенд юзеру
					if err == nil {	// если нет ошибки то отправляем сообщение клиенту о том что удалось поставить юнита что бы он отыграл анимацию.
						// тебе тут надо будет переделать отправку сообщения на отправку пути обьект "TruePatchNode" находиться в проекте по пути " src/mechanics/Phases/movePhase/moveUnit.go "
						ws.WriteJSON(PlaceUnit{Event: "PlaceUnit", Unit: storageUnit}) // todo тут есть ошибка игрок не получает новую зону видимости тут, он это делает в методе с ошибкой, но это не твоя задача
						UpdatePlaceHostilePlayers(actionGame, msg.X, msg.Y)
						return
					} else {
						ws.WriteJSON(ErrorMessage{Event: "Error", Error: "add to db"})
					}
				} else {
					ws.WriteJSON(ErrorMessage{Event: "Error", Error: "place is busy"})
				}
			} else {
				ws.WriteJSON(ErrorMessage{Event: "Error", Error: "place is not allow"})
			}
		}
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "you ready"})
	}
}

func UpdatePlaceHostilePlayers(actionGame *localGame.Game, x, y int) {
	for _, player := range actionGame.GetPlayers() { // смотрим игроков которые участвую в игре

		_, find := player.GetWatchCoordinate(x, y)   // если игрок видит ту зону куда ставиться юнит то отправляем ему сообщение

		if find {
			updater := watchZone.UpdateWatchZone(actionGame, player) // тут мы берем новузю зону видимости т.к.
			// клиент должен увидить юнита которого поставили на карту
			// todo тут есть ошибка т.к. если враг видит юнита то он получит координаты которые видит юнит хотя он не его, но это не твоя задача
			watchPipe <- Watch{Event: "UpdateWatchMap", UserName: player.GetLogin(), GameID: actionGame.Id, Update: updater}
		}
	}
}

type Watch struct {
	Event    string                      `json:"event"`
	UserName string                      `json:"user_name"`
	GameID   int                         `json:"game_id"`
	Update   *watchZone.UpdaterWatchZone `json:"update"`
}

type PlaceUnit struct {
	Event string     `json:"event"`
	Unit  *unit.Unit `json:"unit"`
}
