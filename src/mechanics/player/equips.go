package player

import "../equip"

func (client *Player) SetEquip(equips []*equip.Equip) {
	client.equips = equips
}

func (client *Player) GetEquip() []*equip.Equip {
	return client.equips
}
