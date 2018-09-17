package coordinate

func GetCoordinatesRadius(center *Coordinate, Radius int) []*Coordinate {
	var coordinates = make([]*Coordinate, 0)

	for x := center.X - Radius; x <= center.X+Radius; x++ {
		for y := center.Y - Radius; y <= center.Y+Radius; y++ {
			for z := center.Z - Radius; z <= center.Z+Radius; z++ {
				if x+y+z == 0 {
					coordinates = append(coordinates, &Coordinate{X: x, Y: y, Z: z, Q: x + (z-(z&1))/2, R: z})
				}
			}
		}
	}

	return coordinates
}
