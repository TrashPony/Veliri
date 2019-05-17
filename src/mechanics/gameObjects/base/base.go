package base

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"sync"
	"time"
)

type Base struct {
	ID                        int                      `json:"id"`
	Name                      string                   `json:"name"`
	Q                         int                      `json:"q"`
	R                         int                      `json:"r"`
	MapID                     int                      `json:"map_id"`
	Transports                map[int]*Transport       `json:"transports"`
	Defenders                 map[int]*Transport       `json:"defenders"`
	GravityRadius             int                      `json:"gravity_radius"`
	Respawns                  []*coordinate.Coordinate `json:"respawns"`
	RespawnLock               sync.Mutex               `json:"-"`
	BoundaryAmountOfResources int                      `json:"boundary_amount_of_resources"`
	SumWorkResources          int                      `json:"sum_work_resources"`
	CurrentResources          map[int]*inventory.Slot  `json:"current_resources"` // [id_recycled_type]count
}

type Transport struct {
	ID      int  `json:"id"`
	X       int  `json:"x"`
	Y       int  `json:"y"`
	Job     bool `json:"job"`      /* на задание он или нет */
	Down    bool `json:"down"`     /* на земле он или нет */
	SquadID bool `json:"squad_id"` /* ид того кого он тащит */
}

func (b *Base) CreateTransports(count int) {
	b.Transports = make(map[int]*Transport)

	for i := 0; i < count; i++ {
		b.Transports[i] = &Transport{ID: i, Down: true}
	}
}

func (b *Base) GetFreeTransport() *Transport {
	for _, transport := range b.Transports {
		if !transport.Job {
			return transport
		}
	}
	return nil
}

type defender struct {
	// TODO
}

func (b *Base) CreateDefenders(count int) {
	for i := 0; i < count; i++ {
		// TODO
	}
}

func (b *Base) GetRecyclePercent(resource int) int {

	resourceSlot := b.CurrentResources[resource]

	// при полном достатке ресурса на базе налог будет 50%
	if resourceSlot.Quantity >= b.BoundaryAmountOfResources {
		return 50
	} else {
		// если ресурса нехватает или его вообще нет то налог понижается
		if resourceSlot.Quantity > 0 {
			return resourceSlot.Quantity * 100 / b.BoundaryAmountOfResources
		} else {
			return 0
		}
	}
}

func (b *Base) GetSumEfficiency() int {
	// 0 максимально эффективная база
	// 100 база не функционирует

	percent := 0
	countAllResource := 0

	// если какойто 1 ресурс будет на нуле это дополнительно дает штраф на 10%
	for _, currentResource := range b.CurrentResources {
		if currentResource.Quantity == 0 {
			percent += 10
		} else {
			countAllResource += currentResource.Quantity
		}
	}

	// Если на базе средний показатель по ресурсам выше b.SumWorkResources то налог на обслуживание не добавляется
	if countAllResource < b.SumWorkResources {
		//иначе каждый % нехвата будут добавлять % неэфективности
		percent += countAllResource * 100 / b.SumWorkResources
	}

	// макс размер штрафа 100%
	if percent > 100 {
		percent = 100
	}

	return percent
}

func (b *Base) ConsumptionBaseResource() {
	for {

		consumption := b.SumWorkResources / 10

		for _, currentResources := range b.CurrentResources {
			if currentResources.Quantity > consumption {
				currentResources.Quantity -= consumption
			} else {
				currentResources.Quantity = 0
			}
		}

		time.Sleep(time.Minute)
	}
}
