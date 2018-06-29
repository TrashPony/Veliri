package inventory

import (
	"github.com/gorilla/websocket"
)

func EquipSquad(ws *websocket.Conn, msg Message)  {
	if msg.Event == "AddEquipment" || msg.Event == "ReplaceEquipment" {
		if usersInventoryWs[ws].GetSquad() != nil {
			//equip := inventory.GetTypeEquip(msg.EquipID)

			if msg.Event == "AddEquipment" {
				//usersInventoryWs[ws].Squad.AddEquip(&equip, msg.EquipSlot)
			} else {
				//usersInventoryWs[ws].Squad.ReplaceEquip(&equip, msg.EquipSlot)
			}

			resp := Response{Event: "UpdateSquad", Squad: usersInventoryWs[ws].GetSquad()}
			ws.WriteJSON(resp)

		} else {
			resp := Response{Event: msg.Event, Error: "No select squad"}
			ws.WriteJSON(resp)
		}
	}

	if msg.Event == "RemoveEquipment" {
		if usersInventoryWs[ws].GetSquad() != nil {
			/*err := usersInventoryWs[ws].Squad.DelEquip(msg.EquipSlot)

			if err == nil {
				resp := Response{Event: msg.Event, Error: "none", UnitSlot: msg.EquipSlot}
				ws.WriteJSON(resp)
			} else {
				resp := Response{Event: msg.Event, Error: err.Error(), UnitSlot: msg.EquipSlot}
				ws.WriteJSON(resp)
			}

			resp := Response{Event: "UpdateSquad", Squad: usersInventoryWs[ws].Squad}
			ws.WriteJSON(resp)
		} else {
			resp := Response{Event: msg.Event, Error: "No select squad"}
			ws.WriteJSON(resp)*/
		}
	}
}
