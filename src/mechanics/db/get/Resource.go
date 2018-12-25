package get

import (
	"../../../dbConnect"
	"../../gameObjects/resource"
	"log"
)

func ResourceType() map[int]resource.Resource {
	rows, err := dbConnect.GetDBConnect().Query("SELECT id, name, size, enriched_thorium FROM resource_type")
	if err != nil {
		log.Fatal("get all resource type " + err.Error())
	}
	defer rows.Close()

	allType := make(map[int]resource.Resource)

	for rows.Next() {
		var resourceType resource.Resource

		err := rows.Scan(&resourceType.TypeID, &resourceType.Name, &resourceType.Size,  &resourceType.EnrichedThorium)
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
