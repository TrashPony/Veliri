package globalGame

import (
	"../factories/bases"
	"../gameObjects/base"
	"../gameObjects/detail"
	"../gameObjects/map"
	"../player"
	"errors"
	"github.com/getlantern/deepcopy"
	"github.com/gorilla/websocket"
	"math"
	"sync"
)

const HexagonHeight = 111 // Константы описывающие свойства гексов на игровом поле
const HexagonWidth = 100
const VerticalOffset = HexagonHeight * 3 / 4
const HorizontalOffset = HexagonWidth

type PathUnit struct {
	X           int `json:"x"`
	Y           int `json:"y"`
	Q           int `json:"q"`
	R           int `json:"r"`
	Rotate      int `json:"rotate"`
	Millisecond int `json:"millisecond"`
	Speed       float64
}

func MoveSquad(user *player.Player, ToX, ToY float64, mp *_map.Map, users map[*websocket.Conn]*player.Player) ([]PathUnit, error) {
	startX := float64(user.GetSquad().GlobalX)
	startY := float64(user.GetSquad().GlobalY)
	rotate := user.GetSquad().MatherShip.Rotate

	maxSpeed := float64(user.GetSquad().MatherShip.Speed * 3)
	minSpeed := float64(user.GetSquad().MatherShip.Speed)
	speed := float64(user.GetSquad().MatherShip.Speed)

	// если текущая скорость выше стартовой то берем ее
	if float64(user.GetSquad().MatherShip.Speed) < user.GetSquad().CurrentSpeed {
		speed = user.GetSquad().CurrentSpeed
	}

	var fakeThoriumSlots map[int]*detail.ThoriumSlot

	// копируем что бы не произошло вычетание топлива на расчетах
	err := deepcopy.Copy(&fakeThoriumSlots, &user.GetSquad().MatherShip.Body.ThoriumSlots)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	if user.GetSquad().Afterburner { // если форсаж то х2 скорости
		maxSpeed = maxSpeed * 2
		minSpeed = minSpeed * 2
	}

	err, path := MoveTo(startX, startY, maxSpeed, minSpeed, speed, ToX, ToY, rotate, mp, false,
		fakeThoriumSlots, user.GetSquad().Afterburner, users, user)

	return path, err
}

func LaunchEvacuation(user *player.Player, mp *_map.Map) ([]PathUnit, int, *base.Transport, error) {
	mapBases := bases.Bases.GetBasesByMap(mp.Id)
	minDist := 0.0
	evacuationBase := &base.Base{}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	for _, mapBase := range mapBases {

		x, y := GetXYCenterHex(mapBase.Q, mapBase.R)
		dist := GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
		transport := mapBase.GetFreeTransport()

		if (dist < minDist || minDist == 0) && transport != nil {
			minDist = dist
			evacuationBase = mapBase
		}
	}

	if evacuationBase != nil {
		transport := evacuationBase.GetFreeTransport()
		if transport != nil {
			var startX, startY int

			transport.Job = true

			if transport.X == 0 && transport.Y == 0 {
				startX, startY = GetXYCenterHex(evacuationBase.Q, evacuationBase.R)
			} else {
				startX = transport.X
				startY = transport.Y
			}

			_, path := MoveTo(float64(startX), float64(startY), 250, 15, 15,
				float64(user.GetSquad().GlobalX), float64(user.GetSquad().GlobalY), 0, mp,
				true, nil, false, nil, nil)

			return path, evacuationBase.ID, transport, nil
		} else {
			return nil, 0, nil, errors.New("no available transport")
		}
	} else {
		return nil, 0, nil, errors.New("no available base")
	}
}

func ReturnEvacuation(user *player.Player, mp *_map.Map, baseID int) []PathUnit {
	mapBase, _ := bases.Bases.Get(baseID)
	endX, endY := GetXYCenterHex(mapBase.Q, mapBase.R)

	_, path := MoveTo(float64(user.GetSquad().GlobalX), float64(user.GetSquad().GlobalY), 250, 15, 15,
		float64(endX), float64(endY), 0, mp, true, nil, false, nil, nil)
	return path
}

func MoveTo(forecastX, forecastY, maxSpeed, minSpeed, speed, ToX, ToY float64, rotate int,
	mp *_map.Map, ignoreObstacle bool, thoriumSlots map[int]*detail.ThoriumSlot, afterburner bool,
	users map[*websocket.Conn]*player.Player, user *player.Player) (error, []PathUnit) {

	path := make([]PathUnit, 0)

	fullMax := maxSpeed

	diffRotate := 0 // разница между углом цели и носа корпуса
	dist := 900.0

	for {
		forecastQ := 0
		forecastR := 0
		// находим длинную вектора до цели
		dist = GetBetweenDist(int(forecastX), int(forecastY), int(ToX), int(ToY))
		if dist < 10 {
			break
		}

		if len(path)%10 == 0 || len(path) == 0 {
			if thoriumSlots != nil { // высчитывает эффективность топлива которая влияет на скорость

				efficiency := WorkOutThorium(thoriumSlots, afterburner)
				maxSpeed = (fullMax * efficiency) / 100 // высчитываем максимальную скорость по состоянию топлива

				if efficiency == 0 {
					// кончилось топливо совсем, выходим с ошибкой
					return errors.New("not thorium"), path
				}
			}
		}

		minDistRotate := float64(speed) / ((2 * math.Pi) / float64(360/speed)) // TODO не правильно

		if dist > maxSpeed*25 { // TODO не правильно, тут надо расчитать растояние когда надо сбрасывать скорость
			if int(maxSpeed)*10 != int(speed)*10 {
				if maxSpeed > speed {
					if len(path)%2 == 0 {
						speed += minSpeed / 10
					}
				} else {
					speed -= minSpeed / 10
				}
			} else {
				speed = maxSpeed
			}
		} else {
			if minSpeed < speed {
				if len(path)%2 == 0 {
					speed -= minSpeed / 10
				}
			} else {
				speed = minSpeed
			}
		}

		for i := 0; i < int(speed); i++ { // т.к. за 1 учаток пути корпус может повернуться на много градусов тут этот for)
			needRad := math.Atan2(ToY-forecastY, ToX-forecastX)
			needRotate := int(needRad * 180 / 3.14) // находим какой угол необходимо принять телу

			newRotate := RotateUnit(&rotate, &needRotate)

			if rotate >= needRotate {
				diffRotate = rotate - needRotate
			} else {
				diffRotate = needRotate - rotate
			}

			if diffRotate != 0 { // если разница есть то поворачиваем корпус
				rotate += newRotate
				if minSpeed < speed {
					speed -= minSpeed / (10 * speed) // сбрасывает скорость на поворотах
				}
			} else {
				break
			}
		}

		radRotate := float64(rotate) * math.Pi / 180
		stopX := float64(speed) * math.Cos(radRotate) // идем по вектору движения корпуса
		stopY := float64(speed) * math.Sin(radRotate)

		possibleMove, q, r := CheckCollisionsOnStaticMap(int(forecastX+stopX), int(forecastY+stopY), rotate, mp)

		if users != nil {
			possibleMove = possibleMove && CheckCollisionsPlayers(user, int(forecastX+stopX), int(forecastY+stopY), rotate, mp.Id, users)
		}

		if (diffRotate == 0 || dist > minDistRotate) && (possibleMove || ignoreObstacle) {
			forecastX = forecastX + stopX
			forecastY = forecastY + stopY

			forecastQ = q
			forecastR = r
		} else {
			if diffRotate == 0 {
				break
			}
		}

		path = append(path, PathUnit{X: int(forecastX), Y: int(forecastY), Rotate: rotate, Millisecond: 100,
			Q: forecastQ, R: forecastR, Speed: speed})
	}

	if len(path) > 1 {
		path[len(path)-1].Speed = 0 // на последней точке машина останавливается
	}

	return nil, path
}

func WorkOutThorium(thoriumSlots map[int]*detail.ThoriumSlot, afterburner bool) float64 {
	fullCount := 0

	for _, slot := range thoriumSlots {
		if slot.Count > 0 {
			fullCount++
		}
	}

	efficiency := float64((fullCount * 100) / len(thoriumSlots))

	for _, slot := range thoriumSlots {
		if slot.Count > 0 && efficiency > 0 {

			// формула выроботки топлива, если работает только 1 ячейка из 3х то ее эффективность в 66% больше
			thorium := 1 / float32(100/efficiency)

			if afterburner { // если активирован форсаж то топливо тратиться х5
				thorium = thorium * 15
			}

			slot.WorkedOut += thorium

			if slot.WorkedOut >= 100 {
				slot.Count--
				slot.WorkedOut = 0
			}
		}
	}

	return efficiency
}
