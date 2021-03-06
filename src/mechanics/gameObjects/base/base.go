package base

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"log"
	"strconv"
	"sync"
	"time"
)

type Base struct {
	ID                        int                      `json:"id"`
	Name                      string                   `json:"name"`
	X                         int                      `json:"x"`
	Y                         int                      `json:"y"`
	MapID                     int                      `json:"map_id"`
	Transports                map[int]*Transport       `json:"transports"`
	Defenders                 map[int]*Transport       `json:"defenders"`
	GravityRadius             int                      `json:"gravity_radius"`
	Respawns                  []*coordinate.Coordinate `json:"respawns"`
	RespawnLock               sync.Mutex               `json:"-"`
	BoundaryAmountOfResources int                      `json:"boundary_amount_of_resources"`
	SumWorkResources          int                      `json:"sum_work_resources"`
	CurrentResources          map[int]*inventory.Slot  `json:"current_resources"` // [id_recycled_type]count
	Efficiency                int                      `json:"efficiency"`
	Fraction                  string                   `json:"fraction"`
	Capital                   bool                     `json:"capital"`
}

type Transport struct {
	ID       int     `json:"id"`
	BaseID   int     `json:"base_id"`
	X        int     `json:"x"`
	Y        int     `json:"y"`
	Rotate   int     `json:"rotate"`
	Speed    float64 `json:"speed"`
	Fraction string  `json:"fraction"`
	Job      bool    `json:"job"`      /* на задание он или нет */
	Down     bool    `json:"down"`     /* на земле он или нет */
	SquadID  bool    `json:"squad_id"` /* ид того кого он тащит */
}

func (b *Base) CreateTransports(count int) {
	b.Transports = make(map[int]*Transport)

	for i := 0; i < count; i++ {

		// ид формирует таким тупым способом что бы радар не умер ¯\_(ツ)_/¯,
		// в общем так он будет уникален для каждого транспорта в игре
		idString := strconv.Itoa(b.ID) + strconv.Itoa(i)
		id, err := strconv.Atoi(idString)
		if err != nil {
			log.Fatal(err)
		}

		b.Transports[id] = &Transport{ID: id, Down: true, Fraction: b.Fraction, BaseID: b.ID}
		if b.Transports[id].Fraction == "" {
			b.Transports[id].Fraction = "Replics"
		}
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

func (b *Base) GetRecyclePercent(resource int) int {

	resourceSlot := b.CurrentResources[resource]

	if resourceSlot == nil {
		return 0
	}

	resourceSlot.Tax = 0

	// при полном достатке ресурса на базе налог будет 50%
	if resourceSlot.Quantity >= b.BoundaryAmountOfResources {
		resourceSlot.Tax = 50
	} else {
		// если ресурса нехватает или его вообще нет то налог понижается
		if resourceSlot.Quantity > 0 {
			resourceSlot.Tax = resourceSlot.Quantity * 100 / b.BoundaryAmountOfResources
		}
	}

	// минимальный налог 10%
	if resourceSlot.Tax < 10 {
		resourceSlot.Tax = 10
	}

	return resourceSlot.Tax
}

func (b *Base) GetSumEfficiency() int {
	// 0 максимально эффективная база
	// 100 база не функционирует

	b.Efficiency = 0
	countAllResource := 0

	// если какойто 1 ресурс будет на нуле это дополнительно дает штраф на 10%
	for _, currentResource := range b.CurrentResources {
		if currentResource.Quantity == 0 {
			b.Efficiency += 10
		} else {
			countAllResource += currentResource.Quantity
		}
	}

	// Если на базе средний показатель по ресурсам выше b.SumWorkResources то налог на обслуживание не добавляется
	if countAllResource < b.SumWorkResources {
		//иначе каждый % нехвата будут добавлять % неэфективности
		b.Efficiency += 100 - (countAllResource * 100 / b.SumWorkResources)

		// если все по нулям то база не функционирует
		if countAllResource == 0 {
			b.Efficiency = 100
		}
	}

	// макс размер штрафа 100%
	if b.Efficiency > 100 {
		b.Efficiency = 100
	}

	return b.Efficiency
}

func (b *Base) ConsumptionBaseResource() {
	for {

		consumption := b.SumWorkResources / 500

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
