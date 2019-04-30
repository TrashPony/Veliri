package missions

import (
	missionsDB "github.com/TrashPony/Veliri/src/mechanics/db/missions"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/getlantern/deepcopy"
	"github.com/satori/go.uuid"
	"math/rand"
	"sync"
)

var Missions = NewMissionsStore()

type missions struct {
	missionsType map[int]*mission.Mission
	missions     map[string]*mission.Mission // uuid
	mx           sync.RWMutex
}

func NewMissionsStore() *missions {

	return &missions{
		missionsType: missionsDB.Missions(),
		missions:     make(map[string]*mission.Mission),
	}
}

func (m *missions) GetByID(id int) *mission.Mission {
	return m.missionsType[id]
}

func (m *missions) GetRandomMission() *mission.Mission {
	for {
		// TODO возможны проблемы))
		count := 0
		count2 := rand.Intn(len(m.missionsType))
		for id := range m.missionsType {
			if count == count2 {
				gameMission := m.GetByID(id)
				if gameMission != nil {
					return gameMission
				}
			}
			count++
		}
	}
}

func (m *missions) GenerateMissionForUser() *mission.Mission {

	missionType := m.GetRandomMission()

	if missionType.Name == "Доставка" {
		//в доставке нам нужен в основном только пункт назначения

		uuidNMission := uuid.Must(uuid.NewV4(), nil).String()

		var newMission mission.Mission
		deepcopy.Copy(&newMission, &missionType)

		// назначаем место назначения
		toBase := bases.Bases.GetRandomBase()
		newMission.EndBaseID = toBase.ID
		newMission.ToBase = toBase

		// назначаем и парсим диалоги
		startDialog := gameTypes.Dialogs.GetByID(missionType.StartDialogID)
		startDialog.Mission = uuidNMission
		newMission.StartDialog = startDialog

		endDialog := gameTypes.Dialogs.GetByID(missionType.EndDialogID)
		startDialog.Mission = uuidNMission
		newMission.EndDialog = endDialog

		m.missions[uuidNMission] = &newMission

		return &newMission
	}

	return nil
}
