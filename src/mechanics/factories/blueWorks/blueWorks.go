package blueWorks

import (
	bwDB "../../db/blueWorks"
	"../../gameObjects/blueprints"
)

type blueWorks struct {
	blueWorks map[int]map[int]map[int]*blueprints.BlueWork // [user_ID, base_ID, []works]
}

var BlueWorks = newStorage()

func newStorage() *blueWorks {
	return &blueWorks{
		blueWorks: bwDB.BlueWorks(),
	}
}

func (b *blueWorks) GetAll() map[int]map[int]map[int]*blueprints.BlueWork {
	return b.blueWorks
}

func (b *blueWorks) Add(newWork *blueprints.BlueWork) {
	bwDB.InsertDW(newWork)

	// D8
	if b.blueWorks[newWork.UserID] != nil {
		if b.blueWorks[newWork.UserID][newWork.BaseID] != nil {
			b.blueWorks[newWork.UserID][newWork.BaseID][newWork.ID] = newWork
		} else {
			b.blueWorks[newWork.UserID][newWork.BaseID] = make(map[int]*blueprints.BlueWork)
			b.blueWorks[newWork.UserID][newWork.BaseID][newWork.ID] = newWork
		}
	} else {
		b.blueWorks[newWork.UserID] = make(map[int]map[int]*blueprints.BlueWork)
		b.blueWorks[newWork.UserID][newWork.BaseID] = make(map[int]*blueprints.BlueWork)
		b.blueWorks[newWork.UserID][newWork.BaseID][newWork.ID] = newWork
	}
}

func (b *blueWorks) Remove(work *blueprints.BlueWork) {
	bwDB.Remove(work)
	delete(b.blueWorks[work.UserID][work.BaseID], work.ID)
}
