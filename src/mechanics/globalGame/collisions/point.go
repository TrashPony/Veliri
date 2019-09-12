package collisions

type point struct {
	x int
	y int
}

func (p *point) pointInVector(a, b *point) bool {

	dx1 := int(b.x) - int(a.x)
	dy1 := int(b.y) - int(a.y)

	dx := int(p.x) - int(a.x)
	dy := int(p.y) - int(a.y)

	//вычеслям площадь треуголника, если точка принадлежит вектору то площать будет 0
	s := dx1*dy - dx*dy1

	return s == 0
}
