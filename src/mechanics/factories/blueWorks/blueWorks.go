package blueWorks

import (
	bwDB "../../db/blueWorks"
	"../../gameObjects/blueprints"
	"time"
)

type blueWorks struct {
	blueWorks map[int]*blueprints.BlueWork // id:work
}

var BlueWorks = newStorage()

func newStorage() *blueWorks {
	return &blueWorks{
		blueWorks: bwDB.BlueWorks(),
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

func (b *blueWorks) GetAll() map[int]*blueprints.BlueWork {
	return b.blueWorks
}

func (b *blueWorks) Add(newWork *blueprints.BlueWork) {
	bwDB.InsertDW(newWork)
	b.blueWorks[newWork.ID] = newWork
}

func (b *blueWorks) Remove(removeWork *blueprints.BlueWork) {

	// брать разницу времени текущее и время завершение если оно < 0 то пройтись по всем ордерам этого юзера
	// на этой базе и отнять эту разницу во времени, это для того что бы если ордер отменили что бы время не пропало

	if time.Now().Unix() < removeWork.FinishTime.Unix() {
		deffTime := removeWork.FinishTime.Unix() - time.Now().Unix()

		for _, work := range b.blueWorks {
			if work.UserID == removeWork.ID && work.BaseID == removeWork.BaseID {
				work.FinishTime = time.Unix(work.FinishTime.Unix()-deffTime, 0)
			}
		}
	}

	bwDB.Remove(removeWork)
	delete(b.blueWorks, removeWork.ID)
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
