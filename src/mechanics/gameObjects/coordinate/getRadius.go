package coordinate

func GetCoordinatesRadius(center *Coordinate, Radius int) []*Coordinate {
	var coordinates = make([]*Coordinate, 0)

	// TODO сейчас center.X это col a center.Z это row
	// из за этого не верно работает сеть
	// https://www.redblobgames.com/grids/hexagons/#conversions-offset

	Radius = 1
	for x := center.X - Radius; x <= center.X + Radius; x++ {
		for y := center.Y - Radius; y <= center.Y + Radius; y++ {
			for z := center.Z - Radius; z <= center.Z + Radius; z++ {
				if x + y + z == 0 {
					coordinates = append(coordinates, &Coordinate{X: x + (z - (z&1)) / 2, Y: y, Z: z})
					println(x, y, z)
				}
			}
		}
	}

	return coordinates
}

func GetNeighbors(qCenter, rCenter int) []*Coordinate {

	/*
			// even
			   {-1,0}  {-1,+1}
			{0,-1} {0,0} {0,+1}
			   {+1,0}  {+1,+1}

			// odd
			  {-1,-1}  {-1,0}
			{0,-1} {0,0} {0,+1}
			  {-1,+1}  {+1,0}

	*/

	var coordinates = make([]*Coordinate, 0)

	coordinates = append(coordinates, &Coordinate{X: qCenter, Z: rCenter})

	coordinates = append(coordinates, &Coordinate{X: qCenter - 1, Z: rCenter})
	coordinates = append(coordinates, &Coordinate{X: qCenter, Z: rCenter - 1})
	coordinates = append(coordinates, &Coordinate{X: qCenter + 1, Z: rCenter})
	coordinates = append(coordinates, &Coordinate{X: qCenter, Z: rCenter + 1})

	if rCenter%2 != 0 {
		coordinates = append(coordinates, &Coordinate{X: qCenter + 1, Z: rCenter - 1})
		coordinates = append(coordinates, &Coordinate{X: qCenter + 1, Z: rCenter + 1})
	} else {
		coordinates = append(coordinates, &Coordinate{X: qCenter - 1, Z: rCenter - 1})
		coordinates = append(coordinates, &Coordinate{X: qCenter - 1, Z: rCenter + 1})
	}

	return coordinates
}
