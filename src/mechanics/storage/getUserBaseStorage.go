package storage

import (
	"../db/get"
	inv "../gameObjects/inventory"
	"../player"
)

func GetUserBaseStorage(client *player.Player) *inv.Inventory {
	// TODO брать сторедж из пула, если нет его там то создавать и добавлять в пул
	return get.UserStorage(client.GetID())
}

func AddNewItem(client *player.Player, baseID int, slot *inv.Slot) {
	// TODO сделать пул стореджей для каждого пользователя в видео карты инвентарей, и добавлять/удалять итемы методом
	// TODO инвентаря что бы избежать проблем со слотами
}
