package get

import (
	"../../../dbConnect"
	"../../gameObjects/resource"
	"log"
)

func DetailType() map[int]resource.CraftDetail {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		"id, " +
		"name, " +
		"size, " +
		"enriched_thorium, " +
		"iron, " +
		"copper, " +
		"titanium, " +
		"silicon, " +
		"plastic, " +
		"steel, " +
		"wire " +
		"" +
		" FROM craft_detail")
	if err != nil {
		log.Fatal("get all resource type " + err.Error())
	}
	defer rows.Close()

	allType := make(map[int]resource.CraftDetail)

	for rows.Next() {
		var detailType resource.CraftDetail

		err := rows.Scan(&detailType.TypeID, &detailType.Name, &detailType.Size, &detailType.EnrichedThorium,
			&detailType.Iron, &detailType.Copper, &detailType.Titanium, &detailType.Silicon, &detailType.Plastic,
			&detailType.Steel, &detailType.Wire)
		if err != nil {
			log.Fatal("get all resource type  " + err.Error())
		}

		allType[detailType.TypeID] = detailType
	}

	return allType
}

func ResourceType() map[int]resource.Resource {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		"id, " +
		"name, " +
		"size, " +
		"enriched_thorium, " +
		"iron, " +
		"copper, " +
		"titanium, " +
		"silicon, " +
		"plastic " +
		"" +
		" FROM resource_type")
	if err != nil {
		log.Fatal("get all resource type " + err.Error())
	}
	defer rows.Close()

	allType := make(map[int]resource.Resource)

	for rows.Next() {
		var resourceType resource.Resource

		err := rows.Scan(&resourceType.TypeID, &resourceType.Name, &resourceType.Size, &resourceType.EnrichedThorium,
			&resourceType.Iron, &resourceType.Copper, &resourceType.Titanium, &resourceType.Silicon, &resourceType.Plastic)
		if err != nil {
			log.Fatal("get all resource type  " + err.Error())
		}

		allType[resourceType.TypeID] = resourceType
	}

	return allType
}

func RecycledResourceType() map[int]resource.RecycledResource {
	rows, err := dbConnect.GetDBConnect().Query("SELECT id, name, size FROM recycled_resource_type")
	if err != nil {
		log.Fatal("get all recycled resource type " + err.Error())
	}
	defer rows.Close()

	allType := make(map[int]resource.RecycledResource)

	for rows.Next() {
		var resourceType resource.RecycledResource

		err := rows.Scan(&resourceType.TypeID, &resourceType.Name, &resourceType.Size)
		if err != nil {
			log.Fatal("get all recycled resource type  " + err.Error())
		}

		allType[resourceType.TypeID] = resourceType
	}

	return allType
}

func ReservoirMapType() map[int]resource.Map {

	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		"mtr.id, " +
		"mtr.name, " +
		"mtr.type, " +
		"mtrc.id_base_resource, " +
		"mtrc.max_count, " +
		"mtrc.min_count " +
		" " +
		"FROM map_type_resource mtr, map_type_resource_count mtrc " +
		"WHERE mtr.id = mtrc.id")
	if err != nil {
		log.Fatal("get all map resource type " + err.Error())
	}
	defer rows.Close()

	allType := make(map[int]resource.Map)

	for rows.Next() {
		var reservoirTypeMap resource.Map

		err := rows.Scan(&reservoirTypeMap.TypeID, &reservoirTypeMap.Name, &reservoirTypeMap.Type, &reservoirTypeMap.ResourceID,
			&reservoirTypeMap.MaxCount, &reservoirTypeMap.MinCount)
		if err != nil {
			log.Fatal("get all recycled resource type  " + err.Error())
		}

		allType[reservoirTypeMap.TypeID] = reservoirTypeMap
	}

	return allType
}
