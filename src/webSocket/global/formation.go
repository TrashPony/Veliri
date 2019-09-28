package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/find_path"
)

func NewFormationPos(user *player.Player, msg Message) {

	moveUnit := user.GetSquad().GetUnitByID(msg.UnitID)
	mp, _ := maps.Maps.GetByID(moveUnit.MapID)
	units := globalGame.Clients.GetAllShortUnits(moveUnit.MapID, true)

	if moveUnit.FormationPos == nil {
		x, y, _ := find_path.SearchEndPoint(float64(moveUnit.X), float64(moveUnit.Y), float64(user.GetSquad().MatherShip.X),
			float64(user.GetSquad().MatherShip.Y), moveUnit, mp, units)

		//берем смещение координаты от мс
		moveUnit.FormationPos = &coordinate.Coordinate{
			X: int(x - float64(user.GetSquad().MatherShip.X)),
			Y: int(y - float64(user.GetSquad().MatherShip.Y)),
		}

	} else {
		// TODO позиция не имеет колизии с позициями других юнитов
		moveUnit.FormationPos = &coordinate.Coordinate{X: int(msg.X), Y: int(msg.Y)}
	}

	x, y := user.GetSquad().GetFormationCoordinate(moveUnit.FormationPos.X, moveUnit.FormationPos.Y)
	msg.ToX, msg.ToY = float64(x), float64(y)

	msg.UnitsID = []int{moveUnit.ID}

	Move(user, msg, true)

	go SendMessage(Message{Event: "NewFormationPos", Squad: user.GetSquad(), IDUserSend: user.GetID(), IDMap: moveUnit.MapID, Bot: user.Bot})
}
