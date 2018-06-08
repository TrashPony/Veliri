package player

import "../equip"

func (client *Player) SetEquip(equips []*equip.Equip) {
	client.equips = equips
}

func (client *Player) GetEquips() []*equip.Equip {
	return client.equips
}

func (client *Player) GetEquipByID(id int) (*equip.Equip, bool) {

	for _, playerEquip := range client.equips {
		if playerEquip.Id == id {
			return playerEquip, true
		}
	}

	return nil, false
}
