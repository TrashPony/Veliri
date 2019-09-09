package ai

import (
	"github.com/AvraamMavridis/randomcolor"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/find_path"
	wsGlobal "github.com/TrashPony/Veliri/src/webSocket/global"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

const RespBots = 0

func InitAI() {
	allMaps := maps.Maps.GetAllMap()
	for _, mp := range allMaps {
		mapBases := bases.Bases.GetBasesByMap(mp.Id)
		for _, mapBase := range mapBases {
			for i := 0; i < RespBots; i++ {
				go respBot(mapBase, mp)
				time.Sleep(1 * time.Second)
			}
		}
	}
}

// создаем бота как игрока с отрядом и эквипом
func respBot(base *base.Base, mp *_map.Map) {

	uuidString := uuid.Must(uuid.NewV4(), nil).String()

	newBot := player.Player{
		Login:     uuidString,
		InBaseID:  base.ID,
		Bot:       true,
		UUID:      uuidString,
		Title:     "Вечно зависающий",
		Biography: "Цель моей жизни зависать на точке телепорта",
	}

	body := gameTypes.Bodies.GetRandom()

	// заряжаем топливом
	for _, slot := range body.ThoriumSlots {
		slot.Count = slot.MaxCount
	}

	botSquad := squad.Squad{
		MatherShip: &unit.Unit{
			Body:         body,
			MS:           true,
			HP:           body.MaxHP,
			Power:        body.MaxPower,
			BodyColor1:   "0x" + strings.Split(randomcolor.GetRandomColorInHex(), "#")[1],
			BodyColor2:   "0x" + strings.Split(randomcolor.GetRandomColorInHex(), "#")[1],
			WeaponColor1: "0x" + strings.Split(randomcolor.GetRandomColorInHex(), "#")[1],
			WeaponColor2: "0x" + strings.Split(randomcolor.GetRandomColorInHex(), "#")[1],
			MapID:        mp.Id,
		},
	}

	// ставим рандомную пуху
	weapon := gameTypes.Weapons.GetRandom()

	if botSquad.MatherShip.GetWeaponSlot() != nil {
		botSquad.MatherShip.SetWeaponSlot(
			&detail.BodyWeaponSlot{
				Weapon:     weapon,
				Type:       3,
				Number:     botSquad.MatherShip.GetWeaponSlot().Number,
				WeaponType: weapon.Type,
				HP:         weapon.MaxHP,
			},
		)
	}

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

	respCoordinate := wsGlobal.OutBase(base)

	if respCoordinate == nil {
		return
	}

	x, y := globalGame.GetXYCenterHex(respCoordinate.Q, respCoordinate.R)

	bot.GetSquad().MatherShip.Q = respCoordinate.Q
	bot.GetSquad().MatherShip.R = respCoordinate.R
	bot.GetSquad().MatherShip.Rotate = respCoordinate.RespRotate
	bot.GetSquad().MatherShip.MapID = base.MapID
	bot.GetSquad().MatherShip.X = x
	bot.GetSquad().MatherShip.Y = y
	bot.GetSquad().MatherShip.Evacuation = false
	bot.GetSquad().MatherShip.PointsPath = nil

	// если мы зашли на базу сбрасываем путь бота в ноль
	bot.GlobalPath = nil
	bot.CurrentPoint = 0

	bot.InBaseID = 0

	//оповещаем игроков что бот в игре
	wsGlobal.LoadGame(bot, wsGlobal.Message{})
}

func Transport(bot *player.Player) {

	//-- монитор топлива
	// следим что бы у ботов всегда осталавалось топливо
	go func() {
		for {
			if bot != nil && bot.GetSquad() != nil {
				for _, slot := range bot.GetSquad().MatherShip.Body.ThoriumSlots {
					slot.Count = slot.MaxCount
				}
			}
			time.Sleep(60 * time.Second)
		}
	}()

	//-- транспортник, поиску пути
	for {

		if bot.InBaseID > 0 {
			time.Sleep(15000 * time.Millisecond)
			botBase, _ := bases.Bases.Get(bot.InBaseID)
			outBase(bot, botBase)
		}

		if bot != nil && bot.GetSquad() != nil && !bot.GetSquad().MatherShip.Evacuation && bot.InBaseID == 0 { // todo и есть топливо

			mp, _ := maps.Maps.GetByID(bot.GetSquad().MatherShip.MapID)

			if bot.GetSquad().MatherShip.PointsPath != nil && 0 < len(bot.GetSquad().MatherShip.PointsPath) {

				toPoint := bot.GetSquad().MatherShip.PointsPath[0]
				bot.GetSquad().MatherShip.PointsPath = bot.GetSquad().MatherShip.PointsPath[1:]

				wsGlobal.Move(bot, wsGlobal.Message{ToX: float64(toPoint.X), ToY: float64(toPoint.Y), UnitsID: []int{bot.GetSquad().MatherShip.ID}}, true)

				for {
					time.Sleep(100 * time.Millisecond)
					if !bot.GetSquad().MatherShip.MoveChecker {
						// бот перестал идти
						break
					}

					if mp.Id != bot.GetSquad().MatherShip.MapID {
						//это означает что сектор сменился
						break
					}
				}
			} else {
				getPathAI(bot, mp)
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func getPathAI(bot *player.Player, mp *_map.Map) {
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
			return
		}

		if randMap.Id == bot.GetSquad().MatherShip.MapID {
			toX, toY = globalGame.GetXYCenterHex(randEntryBase.Q, randEntryBase.R)
		} else {
			_, transitionPoints := maps.Maps.FindGlobalPath(bot.GetSquad().MatherShip.MapID, randMap.Id)
			if len(transitionPoints) > 0 {
				// добавляем координату входа в базу в путь
				baseCoordinate, _ := randMap.GetCoordinate(randEntryBase.Q, randEntryBase.R)
				transitionPoints = append(transitionPoints, baseCoordinate)
				bot.GlobalPath = transitionPoints

				bot.CurrentPoint = 0
				toX, toY = globalGame.GetXYCenterHex(transitionPoints[0].Q, transitionPoints[0].R)
			} else {
				return
			}
		}

		// todo преследование игрока

		//рандом
		//xSize, ySize := mp.SetXYSize(globalGame.HexagonWidth, globalGame.HexagonHeight, 1)
		//toX, toY = rand.Intn(xSize), rand.Intn(ySize)

	} else {
		if len(bot.GlobalPath) > bot.CurrentPoint && bot.GlobalPath[bot.CurrentPoint].MapID == bot.GetSquad().MatherShip.MapID {
			toX, toY = globalGame.GetXYCenterHex(bot.GlobalPath[bot.CurrentPoint].Q, bot.GlobalPath[bot.CurrentPoint].R)
		} else {
			bot.GlobalPath = nil
		}
	}

	//println("я иду в х:", toX, " y:", toY)

	// проверка на то что х, у достижимы
	possible, _, _, _ := globalGame.CheckCollisionsOnStaticMap(toX, toY, 0, mp, bot.GetSquad().MatherShip.Body, true)
	if possible {

		go func() {
			for bot.InBaseID == 0 {
				bot.GetSquad().MatherShip.PointsPath = aiSearchPath(toX, toY, bot.GetSquad().MatherShip.X, bot.GetSquad().MatherShip.Y, 50, bot, mp)
				time.Sleep(15 * time.Second)
			}
		}()

	} else {
		println("достижимость: ", possible, bot.GetSquad().MatherShip.MapID, toX, toY)
		bot.GlobalPath = nil
	}
}

func aiSearchPath(toX, toY, startX, startY, scale int, bot *player.Player, mp *_map.Map) []*coordinate.Coordinate {

	if scale < 10 {
		return nil
	}

	mp.SetXYSize(globalGame.HexagonWidth, globalGame.HexagonHeight, scale)

	allUnits := globalGame.Clients.GetAllShortUnits(mp.Id, true)

	_, path := find_path.FindPath(mp, &coordinate.Coordinate{X: startX, Y: startY},
		&coordinate.Coordinate{X: toX, Y: toY}, bot.GetSquad().MatherShip, scale, allUnits)

	if len(path) == 0 {
		return aiSearchPath(toX, toY, startX, startY, scale-10, bot, mp)
	} else {
		return path
	}
}
