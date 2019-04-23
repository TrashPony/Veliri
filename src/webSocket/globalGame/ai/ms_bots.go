package ai

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/find_path"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	wsGlobal "github.com/TrashPony/Veliri/src/webSocket/globalGame"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

const RespBots = 2

func InitAI() {
	allMaps := maps.Maps.GetAllMap()
	for _, mp := range allMaps {
		mapBases := bases.Bases.GetBasesByMap(mp.Id)
		for _, mapBase := range mapBases {
			for i := 0; i < RespBots; i++ {
				go respBot(mapBase, mp)
				time.Sleep(5 * time.Second)
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

	newBot.SetSquad(&botSquad)
	newBot.GetSquad().MatherShip.CalculateParams()

	// делени на характеры
	newBot.Behavior = rand.Intn(2)

	newBot.SetFakeWS(&websocket.Conn{})
	globalGame.Clients.AddNewClient(newBot.GetFakeWS(), &newBot)
	outBase(&newBot, base)

	if newBot.Behavior == 1 || newBot.Behavior == 0 {
		go Transport(&newBot)
	}
}

func outBase(bot *player.Player, base *base.Base) {
	for wsGlobal.CheckTransportCoordinate(base.RespQ, base.RespR, 10, 95, base.MapID) {
		// запускаем механизм проверки и эвакуации игрока с респауна))))
		time.Sleep(time.Millisecond * 100)
	}

	x, y := globalGame.GetXYCenterHex(base.RespQ, base.RespR)

	bot.GetSquad().Q = base.RespQ
	bot.GetSquad().R = base.RespR
	bot.GetSquad().MapID = base.MapID
	bot.GetSquad().GlobalX = x
	bot.GetSquad().GlobalY = y
	bot.GetSquad().Evacuation = false

	// если мы зашли на базу сбрасываем путь бота в ноль
	bot.GlobalPath = nil
	bot.CurrentPoint = 0

	bot.InBaseID = 0
	//оповещаем игроков что бот в игре
	wsGlobal.LoadGame(bot.GetFakeWS(), wsGlobal.Message{})
	// todo после выхода из базы сваливать по прямой быстро а не стоять на токе и распа и строить путь

	// следим что бы у ботов осталавалось топливо
	for _, slot := range bot.GetSquad().MatherShip.Body.ThoriumSlots {
		slot.Count = slot.MaxCount
	}
}

func Transport(bot *player.Player) {
	//-- транспортник, поиску пути
	for {
		if bot.InBaseID > 0 {
			time.Sleep(15000 * time.Millisecond)
			botBase, _ := bases.Bases.Get(bot.InBaseID)
			outBase(bot, botBase)
		}

		if !bot.GetSquad().Evacuation && bot.GetSquad().ActualPath == nil && bot.InBaseID == 0 { // todo и есть топливо
			mp, _ := maps.Maps.GetByID(bot.GetSquad().MapID)
			path := getPathAI(bot, mp)

			//countPossible := 1
			exit := false

			for i := 0; path != nil && i < len(path); i++ {
				if exit {
					bot.GetSquad().ActualPath = nil
					break
				}
				wsGlobal.Move(bot.GetFakeWS(), wsGlobal.Message{ToX: float64(path[i].X), ToY: float64(path[i].Y)})
				for {
					time.Sleep(100 * time.Millisecond)

					if mp.Id != bot.GetSquad().MapID {
						//это означает что сектор сменился
						exit = true
						// если у бота есть цель то она выросла на 1 пройденный сектор)
						bot.CurrentPoint++
						bot.GetSquad().ActualPath = nil
						break
					}

					if !bot.GetSquad().MoveChecker {
						bot.GetSquad().ActualPath = nil
						break
					}
				}
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func getPathAI(bot *player.Player, mp *_map.Map) []*coordinate.Coordinate {
	var toX, toY int

	if bot.GlobalPath == nil {
		// имитация бурной деятельности)

		// сектор
		//toSector := mp.GetEntryTySector(2)
		//if toSector == nil {
		//	return nil
		//}
		//toX, toY = globalGame.GetXYCenterHex(toSector.Q, toSector.R)

		// база
		//toBase := mp.GetRandomEntryBase()
		//toX, toY = globalGame.GetXYCenterHex(toBase.Q, toBase.R)

		// база в другом секторе
		randMap := maps.Maps.GetRandomMap()
		randEntryBase := randMap.GetRandomEntryBase()

		if randEntryBase == nil || randEntryBase.ID == 0 {
			return nil
		}

		if randMap.Id == bot.GetSquad().MapID {
			toX, toY = globalGame.GetXYCenterHex(randEntryBase.Q, randEntryBase.R)
		} else {
			_, transitionPoints := maps.Maps.FindGlobalPath(bot.GetSquad().MapID, randMap.Id)
			if len(transitionPoints) > 0 {
				// добавляем координату входа в базу в путь
				baseCoordinate, _ := randMap.GetCoordinate(randEntryBase.Q, randEntryBase.R)
				transitionPoints = append(transitionPoints, baseCoordinate)
				bot.GlobalPath = transitionPoints

				bot.CurrentPoint = 0
				toX, toY = globalGame.GetXYCenterHex(transitionPoints[0].Q, transitionPoints[0].R)
			} else {
				return nil
			}
		}

		// todo преследование игрока

		//рандом
		//xSize, ySize := mp.SetXYSize(globalGame.HexagonWidth, globalGame.HexagonHeight, 1)
		//toX, toY = rand.Intn(xSize), rand.Intn(ySize)

	} else {
		if len(bot.GlobalPath) > bot.CurrentPoint && bot.GlobalPath[bot.CurrentPoint].MapID == bot.GetSquad().MapID {
			toX, toY = globalGame.GetXYCenterHex(bot.GlobalPath[bot.CurrentPoint].Q, bot.GlobalPath[bot.CurrentPoint].R)
		} else {
			bot.GlobalPath = nil
			return nil
		}
	}

	//println("я иду в х:", toX, " y:", toY)

	// проверка на то что х, у достижимы
	possible, _, _, _ := globalGame.CheckCollisionsOnStaticMap(toX, toY, 0, mp, bot.GetSquad().MatherShip.Body, true)
	if possible {
		path := aiSearchPath(toX, toY, bot.GetSquad().GlobalX, bot.GetSquad().GlobalY, 50, bot, mp)
		return path
	} else {
		println("достижимость: ", possible, bot.GetSquad().MapID, toX, toY)
		bot.GlobalPath = nil
		return nil
	}
}

func aiSearchPath(toX, toY, startX, startY, scale int, bot *player.Player, mp *_map.Map) []*coordinate.Coordinate {

	if scale < 10 {
		return nil
	}

	mp.SetXYSize(globalGame.HexagonWidth, globalGame.HexagonHeight, scale)

	_, path := find_path.FindPath(bot, mp, &coordinate.Coordinate{X: startX, Y: startY},
		&coordinate.Coordinate{X: toX, Y: toY}, bot.GetSquad().MatherShip, scale)

	if len(path) == 0 {
		return aiSearchPath(toX, toY, startX, startY, scale-10, bot, mp)
	} else {
		return path
	}
}
