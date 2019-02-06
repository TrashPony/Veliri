package get

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/blueprints"
	"log"
)

func BlueprintsType() map[int]blueprints.Blueprint {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" name," +
		" item_type," +
		" item_id," +
		" icon," +
		" craft_time," +
		" original," +
		" copies," +
		" enriched_thorium, " +
		" iron, " +
		" copper, " +
		" titanium, " +
		" silicon, " +
		" plastic, " +
		" steel, " +
		" wire," +
		" count," +
		" electronics " +
		"" +
		"FROM blueprints")
	if err != nil {
		log.Fatal("get all blueprints type " + err.Error())
	}
	defer rows.Close()

	allType := make(map[int]blueprints.Blueprint)

	for rows.Next() {
		var blueprintType blueprints.Blueprint

		err := rows.Scan(&blueprintType.ID, &blueprintType.Name, &blueprintType.ItemType, &blueprintType.ItemId,
			&blueprintType.Icon, &blueprintType.CraftTime, &blueprintType.Original, &blueprintType.Copies,
			&blueprintType.EnrichedThorium, &blueprintType.Iron, &blueprintType.Copper, &blueprintType.Titanium,
			&blueprintType.Silicon, &blueprintType.Plastic, &blueprintType.Steel, &blueprintType.Wire,
			&blueprintType.Count, &blueprintType.Electronics)
		if err != nil {
			log.Fatal("get scan all blueprints type " + err.Error())
		}

		allType[blueprintType.ID] = blueprintType
	}

	return allType
}
