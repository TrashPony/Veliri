package squadInventory

import (
	"../db/squad/update"
	"../player"
	"errors"
)

func ChangeSquad(user *player.Player, SquadID int) error {

	squad := user.GetSquadsByID(SquadID)

	if squad != nil {
		if user.GetSquad() != nil {
			user.GetSquad().Active = false         //  старый отряд делаем не активным
			user.GetSquad().BaseID = user.InBaseID // ид базы где храниться отряд
			update.Squad(user.GetSquad(), true)    // обновляем старый отряд в бд
		}

		squad.Active = true
		update.Squad(squad, true)
		GetInventory(user)

		return nil
	} else {
		return errors.New("no squad")
	}
}

func RenameSquad(user *player.Player, newName string) error {
	// TODO валидация данных
	user.GetSquad().Name = newName
	update.Squad(user.GetSquad(), true)

	return nil
}
