package targetPhase

import (
	"../../unit"
	"../../db"
)

func DefendTarget(gameUnit *unit.Unit) {
	// TODO если у юнита нет цели он должен на себя накладывать защитную ауру
	gameUnit.Target = nil

	db.UpdateUnit(gameUnit)
}
