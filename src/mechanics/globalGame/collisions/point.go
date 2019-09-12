package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

type point struct {
	x float64
	y float64
}

func (p *point) pointInVector(a, b *point) bool {

	dx1 := b.x - a.x
	dy1 := b.y - a.y

	dx := p.x - a.x
	dy := p.y - a.y

	//вычеслям площадь треуголника, если точка принадлежит вектору то площать будет 0
	s := dx1*dy - dx*dy1

	// однако если точка находится в векторе но оч далеко то будет ошибка
	// поэтому надо проверить что бы точка была точно внутри вектора с помощью измерения растоиня
	vectorDist := game_math.GetBetweenDist(int(a.x), int(a.y), int(b.x), int(b.y))
	aDist := game_math.GetBetweenDist(int(a.x), int(a.y), int(p.x), int(p.y))
	bDist := game_math.GetBetweenDist(int(b.x), int(b.y), int(p.x), int(p.y))

	return s <= 0.1 && s >= -0.1 && (vectorDist >= aDist && vectorDist >= bDist)
}
