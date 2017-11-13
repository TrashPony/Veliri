package field

import (
	"log"
	"websocket-master"
	"../../game/objects"
	"strconv"
	"sync"
)

var fieldPipe = make(chan FieldResponse)
var initUnit  = make(chan InitUnit)
var coordiante = make(chan Coordinate)
var usersFieldWs = make(map[*websocket.Conn]*Clients) // тут будут храниться наши подключения

var mutex = &sync.Mutex{}

func AddNewUser(ws *websocket.Conn, login string, id int)  {
	CheckDoubleLogin(login, &usersFieldWs)
	usersFieldWs[ws] = &Clients{Login:login, Id:id} // Регистрируем нового Клиента
	print("WS field Сессия: ") // просто смотрим новое подключение
	print(ws)
	println(" login: " + login + " id: " + strconv.Itoa(id))
	defer ws.Close() // Убедитесь, что мы закрываем соединение, когда функция возвращается (с) гугол мужик
	fieldReader(ws, usersFieldWs)
}

func fieldReader(ws *websocket.Conn, usersFieldWs map[*websocket.Conn]*Clients )  {
	for {
		var msg FieldMessage
		err := ws.ReadJSON(&msg) // Читает новое сообщении как JSON и сопоставляет его с объектом Message
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			DelConn(ws, &usersFieldWs , err)
			break
		}

		if msg.Event == "InitGame" {
			InitGame(msg, ws)
			continue
		}

		if msg.Event == "SelectCoordinateCreate" {
			SelectCoordinateCreate(ws)
			continue
		}

		if msg.Event == "CreateUnit" {
			CreateUnit(msg, ws) // TODO второй игрок не может сразу начать строить юнитов
			continue
		}

		if msg.Event == "MouseOver" {
			MouseOver(msg, ws)
			continue
		}

		if msg.Event == "Ready" {
			//Ready(msg, ws)
			continue
		}
		if msg.Event == "SelectUnit" {
			SelectUnit(msg, ws)
			continue
		}
		if msg.Event == "MoveUnit" {
			MoveUnit(msg, ws)
			continue
		}

		if msg.Event == "TargetUnit" {
			TargetUnit(msg, ws)
			continue
		}

		/*if msg.Event == "getPermittedCoordinates" {
			for _, unit := range usersFieldWs[ws].Units {
				SendWatchCoordinate(ws, unit)
			}
			continue
		}*/
	}
}

func FieldReposeSender() {
	for {
		resp := <- fieldPipe
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.Login == resp.UserName {
				err := ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					ws.Close()
					delete(usersFieldWs, ws)
				}
			}
		}
		mutex.Unlock()
	}
}

func InitUnitSender() {
	for {
		resp := <- initUnit
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.Login == resp.UserName {
				err := ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					ws.Close()
					delete(usersFieldWs, ws)
				}
			}
		}
		mutex.Unlock()
	}
}

func CoordinateSender() {
	for {
		resp := <- coordiante
		mutex.Lock()
		for ws, client := range usersFieldWs {
			if client.Login == resp.UserName {
				err := ws.WriteJSON(resp)
				if err != nil {
					log.Printf("error: %v", err)
					ws.Close()
					delete(usersFieldWs, ws)
				}
			}
		}
		mutex.Unlock()
	}
}

type Clients struct { // структура описывающая клиента ws соеденение
	Login string
	Id int
	Watch map[int]map[int]*objects.Coordinate  // map[X]map[Y]
	Units map[int]map[int]*objects.Unit        // map[X]map[Y]
	HostileUnits map[int]map[int]*objects.Unit // map[X]map[Y]
	Map objects.Map
	Respawn objects.Respawn
	CreateZone map[string]*objects.Coordinate
	GameStat objects.Game
	Players []objects.UserStat
}

func (client *Clients) getAllWatchObject(units map[string]*objects.Unit) {
	for _, unit := range units {

		watchCoordinate, watchUnit, err := PermissionCoordinates(client, unit, units)

		if err != nil {
			continue
		}

		for _, xLine := range watchUnit {
			for _, hostile := range xLine {
				if hostile.NameUser != client.Login {
					client.addHostileUnit(hostile)
				} else {
					client.addUnit(unit)
				}
			}
		}

		for _, coordinate := range watchCoordinate {
			client.addCoordinate(coordinate)
		}
	}
}

func (client *Clients) addCoordinate(coordinate *objects.Coordinate) {
	if client.Watch != nil {
		if client.Watch[coordinate.X] != nil {
			client.Watch[coordinate.X][coordinate.Y] = coordinate
		} else {
			client.Watch[coordinate.X] = make(map[int]*objects.Coordinate)
			client.addCoordinate(coordinate)
		}
	} else {
		client.Watch = make(map[int]map[int]*objects.Coordinate)
		client.addCoordinate(coordinate)
	}
}

func (client *Clients) addUnit(unit *objects.Unit) {
	if client.Units != nil {
		if client.Units[unit.X] != nil {
			client.Units[unit.X][unit.Y] = unit
		} else {
			client.Units[unit.X] = make(map[int]*objects.Unit)
			client.addUnit(unit)
		}
	} else {
		client.Units = make(map[int]map[int]*objects.Unit)
		client.addUnit(unit)
	}
}

func (client *Clients) addHostileUnit(hostile *objects.Unit) {
	if client.HostileUnits != nil {
		if client.HostileUnits[hostile.X] != nil {
			client.HostileUnits[hostile.X][hostile.Y] = hostile
		} else {
			client.HostileUnits[hostile.X] = make(map[int]*objects.Unit)
			client.addHostileUnit(hostile)
		}
	} else {
		client.HostileUnits = make(map[int]map[int]*objects.Unit)
		client.addHostileUnit(hostile)
	}
}
