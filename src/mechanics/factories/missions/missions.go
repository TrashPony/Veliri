package missions

import (
	missionsDB "github.com/TrashPony/Veliri/src/mechanics/db/missions"
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	inv "github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
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

/**
* Функция генерит из заготовки экземляр для пользователя под конкретным uuid, который указывается в
* диалогах дабы всегда можно было с диалога перейти в задание
**/
func (m *missions) GenerateMissionForUser(client *player.Player, missionType *mission.Mission) *mission.Mission {

	// если игроку удалось взять это задание то даже если в нем указаны эти параметры они будут совпадать
	startBase, _ := bases.Bases.Get(client.InBaseID)
	startMap, _ := maps.Maps.GetByID(startBase.MapID)

	var newMission mission.Mission
	err := deepcopy.Copy(&newMission, &missionType)
	if err != nil {
		println("GenerateMissionForUser: " + err.Error())
	}

	newMission.StartBase = startBase
	newMission.StartMap = startMap.GetShortInfoMap()
	newMission.UUID = uuid.Must(uuid.NewV4(), nil).String()

	var toMap *_map.Map
	var toBase *base.Base

	// ToSectorName берется из первого экшона где он указан, аналогично и ToBase
	// если action.MapID или action.BaseID равны -1 генерить рандомно
	for _, action := range missionType.Actions {

		if action.MapID > 0 {
			mp, _ := maps.Maps.GetByID(action.MapID)
			toMap = mp
		}

		if action.MapID == -1 {
			for toMap == nil || toMap.Id == startMap.Id {
				toMap = maps.Maps.GetRandomMap()
			}
			action.MapID = toMap.Id // докидываем что бы при паринге экшоно не изменить
		}

		if action.BaseID > 0 {
			toBase, _ = bases.Bases.Get(action.BaseID)
		}

		if action.BaseID == -1 {
			for toBase == nil || toBase.ID == startBase.ID {
				toBase = bases.Bases.GetRandomBase()
			}
			action.BaseID = toBase.ID // докидываем что бы при паринге экшоно не изменить
		}
	}

	startDialog := gameTypes.Dialogs.GetByID(missionType.StartDialogID)
	// указываем принадлежность диалога к миссии
	startDialog.Mission = newMission.UUID
	startDialog.ProcessingDialogText(client.GetLogin(), startBase.Name, toBase.Name, toMap.Name, client.Fraction)
	newMission.StartDialog = startDialog

	for _, action := range newMission.Actions {

		// если action.MapID или action.BaseID (или другие мета) равны -1 генерить рандомно
		if action.MapID == -1 {
			action.MapID = maps.Maps.GetRandomMap().Id
		}

		if action.BaseID == -1 {
			action.BaseID = bases.Bases.GetRandomBase().ID
		}

		if action.DialogID > 0 {
			actionDialog := gameTypes.Dialogs.GetByID(action.DialogID)
			// указываем принадлежность диалога к миссии
			actionDialog.Mission = newMission.UUID
			actionDialog.ProcessingDialogText(client.GetLogin(), startBase.Name, toBase.Name, toMap.Name, client.Fraction)
			action.Dialog = actionDialog
		}
	}

	m.missions[newMission.UUID] = &newMission
	return &newMission
}

/**
* после взятия задания игроком, надо запустить воркеры отслеживания действий
**/
func (m *missions) AcceptMission(client *player.Player, uuid string) *mission.Mission {

	acceptMission, ok := m.missions[uuid]
	if ok {
		m.StartWorkersMonitor(client, acceptMission)
		client.Missions[acceptMission.UUID] = acceptMission
		client.NotifyQueue[acceptMission.UUID] = &player.Notify{Name: "mission", UUID: acceptMission.UUID, Event: "new", Data: acceptMission}
		return acceptMission
	}

	dbPlayer.UpdateUser(client)
	return nil
}

/**
* функция запускат воркеры
**/
func (m *missions) StartWorkersMonitor(client *player.Player, gameMission *mission.Mission) {
	for _, action := range gameMission.Actions {

		if action.TypeFuncMonitor == "delivery_item" {

		}
		if action.TypeFuncMonitor == "get_item_on_base" {

		}
		if action.TypeFuncMonitor == "get_item_on_obj" {

		}
		if action.TypeFuncMonitor == "place_item_in_obj" {

		}
		if action.TypeFuncMonitor == "attack_map_obj" {

		}
		if action.TypeFuncMonitor == "to_q_r" {

		}
		if action.TypeFuncMonitor == "talk_with_base" {

		}
		if action.TypeFuncMonitor == "extract_item" {

		}
	}
	//deliveryItem, _ := gameTypes.TrashItems.GetByID(acceptMission.DeliveryItemId)
	//storages.Storages.AddItem(client.GetID(), client.InBaseID, deliveryItem, "trash", deliveryItem.ID,
	//	1, 1, deliveryItem.Size, 1, false)
}
