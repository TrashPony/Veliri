package missions

import (
	missionsDB "github.com/TrashPony/Veliri/src/mechanics/db/missions"
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	inv "github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/getlantern/deepcopy"
	"github.com/satori/go.uuid"
	"math/rand"
	"sync"
)

var Missions = newMissionsStore()

type missions struct {
	missionsType map[int]*mission.Mission
	missions     map[string]*mission.Mission // uuid
	mx           sync.RWMutex
}

func newMissionsStore() *missions {

	return &missions{
		missionsType: missionsDB.Missions(),
		missions:     make(map[string]*mission.Mission),
	}
}

func (m *missions) GetByID(id int) *mission.Mission {
	return m.missionsType[id]
}

func (m *missions) GetStory(episode int, fraction string) *mission.Mission {
	for _, miss := range m.missionsType {
		if miss.MainStory && miss.Story == episode && miss.Fraction == fraction {
			return miss
		}
	}
	return nil
}

func (m *missions) GetAllMissType() map[int]*mission.Mission {
	return m.missionsType
}

func (m *missions) GetRandomMission() *mission.Mission {
	for {
		// TODO возможны проблемы))
		count := 0
		count2 := rand.Intn(len(m.missionsType))
		for id, miss := range m.missionsType {
			if count == count2 && !miss.MainStory {
				gameMission := m.GetByID(id)
				if gameMission != nil {
					return gameMission
				}
			}
			count++
		}
	}
}

func (m *missions) SaveTypeMission(mission *mission.Mission) {
	oldType, _ := m.missionsType[mission.ID]
	missionsDB.UpdateMission(mission, oldType)
	m.missionsType[mission.ID] = mission
}

func (m *missions) DeleteMission(mission *mission.Mission) {
	missionsDB.DeleteMission(mission)
	delete(m.missionsType, mission.ID)
}

func (m *missions) RemoveItem(missionID, slot int) {
	typeMiss, _ := m.missionsType[missionID]
	delete(typeMiss.RewardItems.Slots, slot)
	m.SaveTypeMission(typeMiss)
}

func (m *missions) AddItem(missionID, itemID, itemQuantity int, itemType string) {
	typeMiss, _ := m.missionsType[missionID]

	slotNumber := typeMiss.RewardItems.GetEmptySlot()
	if slotNumber > 0 {
		var inventorySlot = inv.Slot{Type: itemType, ItemID: itemID, Quantity: itemQuantity, PlaceUserID: 0}
		inv.FillSlot(&inventorySlot, slotNumber, typeMiss.RewardItems, true)
	}

	m.SaveTypeMission(typeMiss)
}

func (m *missions) ActionAddItem(actionID, itemID, itemQuantity int, itemType string) {
	for _, miss := range m.missionsType {
		for _, action := range miss.Actions {
			if action.ID == actionID {
				var inventorySlot = inv.Slot{Type: itemType, ItemID: itemID, Quantity: itemQuantity, PlaceUserID: 0}
				slotNumber := action.NeedItems.GetEmptySlot()
				if slotNumber > 0 {
					inv.FillSlot(&inventorySlot, slotNumber, action.NeedItems, true)
					m.SaveTypeMission(miss)
				}
				return
			}
		}
	}
}

func (m *missions) ActionRemoveItem(actionID, slot int) {
	for _, miss := range m.missionsType {
		for _, action := range miss.Actions {
			if action.ID == actionID {
				delete(action.NeedItems.Slots, slot)
				m.SaveTypeMission(miss)
				return
			}
		}
	}
}

func (m *missions) AddMission(name string) {
	newInventory := &inv.Inventory{Slots: make(map[int]*inv.Slot)}
	newInventory.SetSlotsSize(999)

	newMission := mission.Mission{
		Name:        name,
		Actions:     make([]*mission.Action, 0),
		RewardItems: newInventory,
	}
	missionsDB.AddMission(&newMission)
	m.missionsType[newMission.ID] = &newMission
}

func (m *missions) GenerateMissionForUser(client *player.Player) *mission.Mission {

	missionType := m.GetRandomMission()

	startBase, _ := bases.Bases.Get(client.InBaseID)
	startMap, _ := maps.Maps.GetByID(startBase.MapID)

	if missionType.Type == "delivery" {
		//в доставке нам нужен в основном только пункт назначения
		// и там только 1 экшон это соотвественно начать диалог на базе назначения

		var newMission mission.Mission
		deepcopy.Copy(&newMission, &missionType)

		newMission.StartBase = startBase
		newMission.StartMap = startMap.GetShortInfoMap()

		newMission.UUID = uuid.Must(uuid.NewV4(), nil).String()

		// назначаем место назначения, ищем пока не найдет базу которая не база игрока где он берет квест)
		var toBase *base.Base
		for toBase == nil || toBase.ID == startBase.ID {
			toBase = bases.Bases.GetRandomBase()
		}

		toMap, _ := maps.Maps.GetByID(toBase.MapID)

		// назначаем и парсим диалоги
		startDialog := gameTypes.Dialogs.GetByID(missionType.StartDialogID)
		startDialog.Mission = newMission.UUID
		startDialog.ProcessingDialogText(client.GetLogin(), startBase.Name, toBase.Name, toMap.Name, client.Fraction)
		newMission.StartDialog = startDialog

		for _, action := range newMission.Actions {
			endDialog := gameTypes.Dialogs.GetByID(action.DialogID)
			endDialog.Mission = newMission.UUID

			endDialog.ProcessingDialogText(client.GetLogin(), startBase.Name, toBase.Name, toMap.Name, client.Fraction)
			action.Dialog = endDialog
			action.BaseID = toBase.ID
		}

		m.missions[newMission.UUID] = &newMission
		return &newMission
	}
	return nil
}

func (m *missions) AcceptMission(client *player.Player, uuid string) *mission.Mission {

	acceptMission, ok := m.missions[uuid]
	if ok {
		if acceptMission.Type == "delivery" {

			//deliveryItem, _ := gameTypes.TrashItems.GetByID(acceptMission.DeliveryItemId)
			//storages.Storages.AddItem(client.GetID(), client.InBaseID, deliveryItem, "trash", deliveryItem.ID,
			//	1, 1, deliveryItem.Size, 1, false)
			//client.Missions[acceptMission.UUID] = acceptMission

			client.NotifyQueue[acceptMission.UUID] = &player.Notify{Name: "mission", UUID: acceptMission.UUID, Event: "new", Data: acceptMission}
		}
		return acceptMission
	}

	dbPlayer.UpdateUser(client)
	return nil
}
