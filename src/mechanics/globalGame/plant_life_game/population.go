package plant_life_game

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math/rand"
	"time"
)

type lifeObject struct {
	life               bool
	futureLife         bool
	countLifeNeighbors int
	Neighbors          []*coordinate.Coordinate
	X                  int
	Y                  int
}

var typeSources = map[string][]coordinate.Coordinate{
	"glider":     {{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: -1}, {X: 1, Y: -2}},
	"flasher":    {{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}},
	"locomotive": {{X: 0, Y: 0}, {X: -1, Y: -1}, {X: -1, Y: -2}, {X: -1, Y: -3}, {X: -1, Y: -4}, {X: 0, Y: -5}},
	// That'S a penis! ..... :)
	"penis": {{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 1, Y: -1}, {X: 1, Y: -2}, {X: 1, Y: -3}, {X: 1, Y: -4}},
	"space_ship": {
		{X: 0, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: -2}, {X: 1, Y: -3},
		{X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}, {X: 4, Y: 0}, {X: 5, Y: -1},
		{X: 5, Y: -1}, {X: 3, Y: -4},
	},
	"hive": {
		{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1},
		{X: 1, Y: -2}, {X: 1, Y: 2},
		{X: 2, Y: -3}, {X: 2, Y: 3},
		{X: 3, Y: -3}, {X: 3, Y: 3},
		{X: 4, Y: 0},
		{X: 5, Y: -2}, {X: 5, Y: 2},
		{X: 6, Y: -1}, {X: 6, Y: 1}, {X: 6, Y: 0},
		{X: 7, Y: 0},
	},
	"Gosper_glider_gun": {
		// box
		{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 1, Y: -1},
		// hive
		{X: 10, Y: 0}, {X: 10, Y: 1}, {X: 10, Y: -1},
		{X: 11, Y: -2}, {X: 11, Y: 2},
		{X: 12, Y: -3}, {X: 12, Y: 3},
		{X: 13, Y: -3}, {X: 13, Y: 3},
		{X: 14, Y: 0},
		{X: 15, Y: -2}, {X: 15, Y: 2},
		{X: 16, Y: -1}, {X: 16, Y: 1}, {X: 16, Y: 0},
		{X: 17, Y: 0},
		// xz
		{X: 20, Y: -1}, {X: 20, Y: -2}, {X: 20, Y: -3},
		{X: 21, Y: -1}, {X: 21, Y: -2}, {X: 21, Y: -3},
		{X: 22, Y: -4}, {X: 22, Y: 0},
		{X: 24, Y: 1}, {X: 24, Y: 0}, {X: 24, Y: -4}, {X: 24, Y: -5},
		// box 2
		{X: 34, Y: -2}, {X: 34, Y: -3}, {X: 35, Y: -2}, {X: 35, Y: -3},
	},
}

func getRandomSource() []coordinate.Coordinate {
	for {
		count := 0
		count2 := rand.Intn(len(typeSources))
		for id := range typeSources {
			if count == count2 {
				return typeSources[id]
			}
			count++
		}
	}
}

func initPopulationSource(source []coordinate.Coordinate, x, y int, texturePlant string, mp *_map.Map) {
	// TODO поворот на 90, 180 и 270 градусов

	// костыль Х)
	x *= game_math.CellSize
	y *= game_math.CellSize

	for _, offset := range source {
		createPlant(
			x+(offset.X*game_math.CellSize)+game_math.CellSize/2, // game_math.CellSize/2 что бы брался пиксель в центре клетки
			y+(offset.Y*game_math.CellSize)+game_math.CellSize/2,
			texturePlant,
			mp,
		)
	}
}

func PopulationWorker(mp *_map.Map) {
	for {

		// если на карте меньше N кустов то рождаем новые)
		if GetCountPlantIMap("plant_4", mp) < 200 {
			initPopulationSource(
				getRandomSource(),
				rand.Intn(mp.XSize/game_math.CellSize-6)+3, // -6+3 что бы совсем на краях не делались
				rand.Intn(mp.YSize/game_math.CellSize-6)+3,
				"plant_4",
				mp)
		}

		time.Sleep(time.Millisecond * 2000)
	}
}
