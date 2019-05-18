package blueWorks

import (
	bwDB "github.com/TrashPony/Veliri/src/mechanics/db/blueWorks"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/blueprints"
	"sync"
	"time"
)

type blueWorks struct {
	blueWorks map[int]*blueprints.BlueWork // id:work
	mx        sync.Mutex
}

var BlueWorks = newStorage()

func newStorage() *blueWorks {
	works := bwDB.BlueWorks()

	for _, work := range works {
		// докидываем обьекты ворка
		work.Blueprint, _ = gameTypes.BluePrints.GetByID(work.BlueprintID)
		work.Item = gameTypes.BluePrints.GetItemsByBluePrintID(work.BlueprintID)
	}

	return &blueWorks{
		blueWorks: works,
	}
}

func (b *blueWorks) GetByID(id int) *blueprints.BlueWork {
	for _, work := range b.blueWorks {
		if work.ID == id {
			return work
		}
	}
	return nil
}

func (b *blueWorks) GetByUserAndBase(userID, baseID int) map[int]*blueprints.BlueWork {

	works := make(map[int]*blueprints.BlueWork)

	for _, work := range b.blueWorks {
		if work.UserID == userID && work.BaseID == baseID {
			works[work.ID] = work
		}
	}
	return works
}

func (b *blueWorks) GetSameWorks(bpID, MineralTax, TimeTax, userID, baseID int, toTime, startTime int64) map[int]*blueprints.BlueWork {

	// startTime - от какого времени брать
	// toTime - до какого времени брать

	works := make(map[int]*blueprints.BlueWork)

	for _, work := range b.blueWorks {
		if work.UserID == userID && work.BaseID == baseID && work.BlueprintID == bpID &&
			work.MineralTaxPercentage == MineralTax && work.TimeTaxPercentage == TimeTax && work.GetDonePercent() < 0 &&
			work.FinishTime.UTC().Unix() <= toTime && work.FinishTime.UTC().Unix() >= startTime {

			works[work.ID] = work
		}
	}
	return works
}

func (b *blueWorks) GetAll() map[int]*blueprints.BlueWork {
	return b.blueWorks
}

func (b *blueWorks) Add(newWork *blueprints.BlueWork) {

	newWork.Blueprint, _ = gameTypes.BluePrints.GetByID(newWork.BlueprintID)
	newWork.Item = gameTypes.BluePrints.GetItemsByBluePrintID(newWork.BlueprintID)

	bwDB.InsertDW(newWork)
	b.blueWorks[newWork.ID] = newWork
}

func (b *blueWorks) Remove(removeWork *blueprints.BlueWork) {
	b.mx.Lock()
	defer b.mx.Unlock()
	// брать разницу времени текущее и время завершение если оно < 0 то пройтись по всем ордерам этого юзера
	// на этой базе и отнять эту разницу во времени, это для того что бы если ордер отменили что бы время не пропало

	// если заказ еще не делается то надо вычитать время его работы только у тех кто выше по времени

	var deffTime int64

	if removeWork.GetDonePercent() > 0 {
		deffTime = removeWork.FinishTime.Unix() - time.Now().Unix()
	} else {
		deffTime = time.Unix(int64(removeWork.Blueprint.CraftTime+(removeWork.Blueprint.CraftTime*removeWork.TimeTaxPercentage/100)), 0).Unix()
	}

	for _, work := range b.blueWorks {
		if work.UserID == removeWork.UserID && work.BaseID == removeWork.BaseID &&
			removeWork.FinishTime.Unix() < work.FinishTime.Unix() {

			work.FinishTime = time.Unix(work.FinishTime.Unix()-deffTime, 0)
			bwDB.UpdateBW(work)
		}
	}

	delete(b.blueWorks, removeWork.ID)
	bwDB.Remove(removeWork)
}

func (b *blueWorks) GetWorkTime(userID, baseID int) int64 {
	//метод возвраает время когда линия произвосдвта будет свободна
	maxTime := time.Now().Unix()

	for _, work := range b.blueWorks {
		if work.UserID == userID && work.BaseID == baseID && maxTime < work.FinishTime.Unix() {
			maxTime = work.FinishTime.Unix()
		}
	}

	return maxTime
}
