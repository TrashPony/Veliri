package plant_life_game

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/dynamic_objects"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"math/rand"
)

func createPlant(x, y int, texturePlant string, mp *_map.Map) {
	plant := dynamic_objects.DynamicObjects.GetDynamicObjectByTexture(texturePlant, rand.Intn(360))
	plant.X = x
	plant.Y = y
	plant.Scale = 1

	minSize := 5
	// что бы шанс получить большого был не велик
	if rand.Intn(4) == 1 {
		plant.MaxScale = rand.Intn(100-minSize) + minSize
	} else {
		plant.MaxScale = rand.Intn(40-minSize) + minSize
	}

	plant.HP = plant.MaxHP
	plant.CalculateScale()

	if !collisions.CheckObjectCollision(plant, mp, true) {
		mp.AddDynamicObject(plant)
	}
}
