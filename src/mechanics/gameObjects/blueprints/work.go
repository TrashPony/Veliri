package blueprints

import "time"

type BlueWork struct {
	ID                   int         `json:"id"`
	BlueprintID          int         `json:"blueprint_id"`
	BaseID               int         `json:"base_id"`
	UserID               int         `json:"user_id"`
	FinishTime           time.Time   `json:"finish_time"`
	MineralTaxPercentage int         `json:"mineral_tax_percentage"`
	TimeTaxPercentage    int         `json:"time_tax_percentage"`
	Blueprint            *Blueprint  `json:"blueprint"`
	Item                 interface{} `json:"item"`
}

func (w *BlueWork) GetDonePercent() int {
	realTimeCraft := time.Unix(int64(w.Blueprint.CraftTime+(w.Blueprint.CraftTime*w.TimeTaxPercentage/100)), 0)
	startTime := time.Unix(w.FinishTime.UTC().Unix()-realTimeCraft.UTC().Unix(), 0)
	diffTime := time.Unix(time.Now().UTC().Unix()-startTime.UTC().Unix(), 0)
	return int(diffTime.UTC().Unix() * 100 / realTimeCraft.UTC().Unix())
}
