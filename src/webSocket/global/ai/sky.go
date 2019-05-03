package ai

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	wsGlobal "github.com/TrashPony/Veliri/src/webSocket/global"
	"github.com/satori/go.uuid"
	"math"
	"math/rand"
	"strconv"
	"time"
)

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

	Uuid := uuid.Must(uuid.NewV4(), nil)

	randomCloud := "cloud" + strconv.Itoa(rand.Intn(13))
	randomPos := rand.Intn(2)

	sizeMapX := (globalGame.HexagonWidth+5)*mp.QSize + 500
	sizeMapY := 185 * mp.RSize / 2
	speed := rand.Intn(40) + 20
	alpha := 0.2 + rand.Float64()*(0.8-0.2)

	var newCloud *wsGlobal.Cloud
	if randomPos == 0 {
		newCloud = &wsGlobal.Cloud{
			Name:     randomCloud,
			Uuid:     Uuid.String(),
			Speed:    speed,
			Alpha:    alpha,
			SizeMapX: sizeMapX,
			SizeMapY: sizeMapY,
			X:        rand.Intn(sizeMapX),
			Y:        -500,
			Angle:    135,
			IDMap:    mp.Id,
		}
	}
	if randomPos == 1 {
		newCloud = &wsGlobal.Cloud{
			Name:     randomCloud,
			Uuid:     Uuid.String(),
			Speed:    speed,
			Alpha:    alpha,
			SizeMapX: sizeMapX,
			SizeMapY: sizeMapY,
			X:        sizeMapX + 500,
			Y:        rand.Intn(sizeMapY),
			Angle:    135,
			IDMap:    mp.Id,
		}
	}

	go MoveCloud(newCloud)
}

func MoveCloud(cloud *wsGlobal.Cloud) {
	for {
		time.Sleep(1000 * time.Millisecond)

		radRotate := float64(cloud.Angle) * math.Pi / 180
		cloud.X = int(float64(cloud.Speed)*math.Cos(radRotate)) + cloud.X // идем по вектору движения
		cloud.Y = int(float64(cloud.Speed)*math.Sin(radRotate)) + cloud.Y

		go wsGlobal.SendMessage(wsGlobal.Message{Event: "MoveCloud", Cloud: cloud, IDMap: cloud.IDMap})

		if cloud.X > cloud.SizeMapX+500 || cloud.Y > cloud.SizeMapY+500 || -500 > cloud.X || -500 > cloud.Y {
			go CreateCloud(cloud.IDMap)
			break
		}
	}
}
