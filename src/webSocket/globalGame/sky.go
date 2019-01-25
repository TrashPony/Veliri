package globalGame

import (
	"../../mechanics/factories/maps"
	"../../mechanics/globalGame"
	"github.com/satori/go.uuid"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type cloud struct {
	Name     string  `json:"name"`
	Speed    int     `json:"speed"`
	Alpha    float64 `json:"alpha"`
	X        int     `json:"x"`
	Y        int     `json:"y"`
	Angle    int     `json:"angle"`
	Uuid     string  `json:"uuid"`
	sizeMapX int
	sizeMapY int
	idMap    int
}

func SkyGenerator() {
	allMaps := maps.Maps.GetAllMap()
	for _, mp := range allMaps {
		for i := 0; i < 15; i++ {
			CreateCloud(mp.Id)
		}
	}
}

func CreateCloud(mapID int) {

	mp, _ := maps.Maps.GetByID(mapID)

	Uuid := uuid.Must(uuid.NewV4())

	randomCloud := "cloud" + strconv.Itoa(rand.Intn(13))
	randomPos := rand.Intn(2)

	sizeMapX := (globalGame.HexagonWidth+5)*mp.QSize + 500
	sizeMapY := 185 * mp.RSize / 2
	speed := rand.Intn(40) + 20
	alpha := 0.2 + rand.Float64()*(0.8-0.2)

	var newCloud *cloud
	if randomPos == 0 {
		newCloud = &cloud{
			Name:     randomCloud,
			Uuid:     Uuid.String(),
			Speed:    speed,
			Alpha:    alpha,
			sizeMapX: sizeMapX,
			sizeMapY: sizeMapY,
			X:        rand.Intn(sizeMapX),
			Y:        -500,
			Angle:    135,
			idMap:    mp.Id,
		}
	}
	if randomPos == 1 {
		newCloud = &cloud{
			Name:     randomCloud,
			Uuid:     Uuid.String(),
			Speed:    speed,
			Alpha:    alpha,
			sizeMapX: sizeMapX,
			sizeMapY: sizeMapY,
			X:        sizeMapX + 500,
			Y:        rand.Intn(sizeMapY),
			Angle:    135,
			idMap:    mp.Id,
		}
	}

	go MoveCloud(newCloud)
}

func MoveCloud(cloud *cloud) {
	for {
		time.Sleep(1000 * time.Millisecond)

		radRotate := float64(cloud.Angle) * math.Pi / 180
		cloud.X = int(float64(cloud.Speed)*math.Cos(radRotate)) + cloud.X // идем по вектору движения
		cloud.Y = int(float64(cloud.Speed)*math.Sin(radRotate)) + cloud.Y

		globalPipe <- Message{Event: "MoveCloud", Cloud: cloud, idMap: cloud.idMap}

		if cloud.X > cloud.sizeMapX+500 || cloud.Y > cloud.sizeMapY+500 || -500 > cloud.X || -500 > cloud.Y {
			go CreateCloud(cloud.idMap)
			break
		}
	}
}
