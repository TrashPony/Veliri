package blueprints

import "time"

type BlueWork struct {
	ID                      int       `json:"id"`
	BlueprintID             int       `json:"blueprint_id"`
	BaseID                  int       `json:"base_id"`
	UserID                  int       `json:"user_id"`
	FinishTime              time.Time `json:"finish_time"`
	MineralSavingPercentage int       `json:"mineral_saving_percentage"`
	TimeSavingPercentage    int       `json:"time_saving_percentage"`
}
