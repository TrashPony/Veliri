package missions

import (
	missionsDB "github.com/TrashPony/Veliri/src/mechanics/db/missions"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"sync"
)

var Missions = NewMissionsStore()

type missions struct {
	missions map[int]*mission.Mission
	mx       sync.RWMutex
}

func NewMissionsStore() *missions {
	return &missions{
		missions: missionsDB.Missions(),
	}
}
