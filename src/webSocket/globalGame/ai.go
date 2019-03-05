package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/find_path"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

const RespBots = 30

func InitAI() {
	allMaps := maps.Maps.GetAllMap()
	for _, mp := range allMaps {
		mapBases := bases.Bases.GetBasesByMap(mp.Id)
		for _, mapBase := range mapBases {
			for i := 0; i < RespBots; i++ {
				// на карту базу делаем по 3 юнита
				go respBot(mapBase, mp)
				time.Sleep(10 * time.Second)
			}
		}
	}
}

func respBot(base *base.Base, mp *_map.Map) {
	// собираем их тела с рандомным эквипом
	newBot := player.Player{InBaseID: base.ID, Bot: true, UUID: uuid.Must(uuid.NewV4(), nil).String()}

	body := gameTypes.Bodies.GetRandom()

	// заряжаем топливом
	for _, slot := range body.ThoriumSlots {
		slot.Count = slot.MaxCount
	}

	botSquad := squad.Squad{
		MatherShip: &unit.Unit{
			Body:  body,
			MS:    true,
			HP:    body.MaxHP,
			Power: body.MaxPower,
		},
		MapID: mp.Id,
	}

	// ставим рандомную пуху
	weapon := gameTypes.Weapons.GetRandom()

	botSquad.MatherShip.SetWeaponSlot(
		&detail.BodyWeaponSlot{
			Weapon:     weapon,
			Type:       3,
			Number:     3, // TODO проблема, номер может быть разный
			WeaponType: weapon.Type,
			HP:         weapon.MaxHP,
		},
	)

	botSquad.CreateMove()
	newBot.SetSquad(&botSquad)
	newBot.GetSquad().MatherShip.CalculateParams()

	// делени на характеры
	newBot.Behavior = rand.Intn(2)

	newBot.SetFakeWS(&websocket.Conn{})
	Clients.addNewClient(newBot.GetFakeWS(), &newBot)
	outBase(&newBot, base)

	if newBot.Behavior == 1 || newBot.Behavior == 0 {
		go Transport(&newBot, mp)
	}
}

func outBase(bot *player.Player, base *base.Base) {
	for CheckTransportCoordinate(base.RespQ, base.RespR, 10, 75, base.MapID) {
		// запускаем механизм проверки и эвакуации игрока с респауна))))
		time.Sleep(time.Millisecond * 100)
	}

	x, y := globalGame.GetXYCenterHex(base.RespQ, base.RespR)

	bot.GetSquad().Q = base.RespQ
	bot.GetSquad().R = base.RespR
	bot.GetSquad().MapID = base.MapID
	bot.GetSquad().GlobalX = x
	bot.GetSquad().GlobalY = y

	bot.InBaseID = 0
	//оповещаем игроков что бот в игре
	loadGame(bot.GetFakeWS(), Message{})
}

func Transport(bot *player.Player, mp *_map.Map) {
	//-- транспортник, поиску пути
	for {
		if bot.InBaseID > 0 {
			time.Sleep(10000 * time.Millisecond)
			botBase, _ := bases.Bases.Get(bot.InBaseID)
			outBase(bot, botBase)
		}

		toQ, toR := rand.Intn(mp.QSize), rand.Intn(mp.RSize)

		if !bot.GetSquad().MoveChecker {
			_, path := find_path.FindPath(
				bot,
				mp,
				&coordinate.Coordinate{Q: bot.GetSquad().Q, R: bot.GetSquad().R},
				&coordinate.Coordinate{Q: toQ, R: toR},
				bot.GetSquad().MatherShip,
			)

			//countPossible := 1
			for i := 0; i < len(path); i++ {
				toX, toY := globalGame.GetXYCenterHex(path[i].Q, path[i].R)

				// TODO учитывать наличие игроков при построение маршрута.
				// ПРИМЕР ДЛЯ ПОНИМАНИЯ: если между точкой 5 и точкой 10 нет прептявий то идти напрямик
				//fastPath, _ := globalGame.MoveSquad(bot, float64(toX), float64(toY), mp)
				//
				//if len(fastPath) > 10*countPossible && i < len(path)-1 {
				//	countPossible++
				//	continue
				//} else {
				//	countPossible = 1
				move(bot.GetFakeWS(), Message{ToX: float64(toX), ToY: float64(toY)})
				//}

				for {
					time.Sleep(100 * time.Millisecond)
					if !bot.GetSquad().MoveChecker {
						break
					}
				}
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}
