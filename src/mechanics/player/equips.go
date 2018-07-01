package player

import "../gameObjects/equip"

func (client *Player) SetEquip(equips []*equip.Equip) {
	client.equips = equips
}

func (client *Player) GetEquips() []*equip.Equip {
	return client.equips
}

func (client *Player) GetEquipByID(id int) (*equip.Equip, bool) {

	for _, playerEquip := range client.equips {
		if playerEquip.ID == id {
			return playerEquip, true
		}
	}

	return nil, false
}
