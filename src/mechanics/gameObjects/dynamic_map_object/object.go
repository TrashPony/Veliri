package dynamic_map_object

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/effect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/obstacle_point"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"github.com/getlantern/deepcopy"
	"time"
)

type Object struct {
	// везде где есть приставка Type это оригиналдьные данные типа, все остальное сформированые
	ID                  int    `json:"id"`
	TypeID              int    `json:"type_id"`
	Type                string `json:"type"`
	Texture             string `json:"texture"`
	AnimateSpriteSheets string `json:"animate_sprite_sheets"`
	AnimateLoop         bool   `json:"animate_loop"`
	UnitOverlap         bool   `json:"unit_overlap"`
	Rotate              int    `json:"rotate"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Inventory           bool   `json:"inventory"`
	BoxID               int    `json:"box_id"`
	MaxHP               int    `json:"max_hp"`
	TypeMaxHP           int    `json:"type_max_hp"`
	HP                  int    `json:"hp"`
	Scale               int    `json:"scale"`
	MaxScale            int    `json:"max_scale"` // определяется рандомно для растений максимальный размер куста
	GrowTime            int    `json:"grow_time"` // говорит время цикла когда растение росло для гладкой отрисовки
	Shadow              bool   `json:"shadow"`
	AnimationSpeed      int    `json:"animation_speed"`
	Priority            int    `json:"priority"`

	X int `json:"x"`
	Y int `json:"y"`

	TypeXShadowOffset int `json:"type_x_shadow_offset"`
	XShadowOffset     int `json:"x_shadow_offset"`
	TypeYShadowOffset int `json:"type_y_shadow_offset"`
	YShadowOffset     int `json:"y_shadow_offset"`
	ShadowIntensity   int `json:"shadow_intensity"`

	Dialog *dialog.Dialog `json:"dialog"`

	Destroyed   bool      `json:"destroyed"`
	DestroyTime time.Time `json:"destroy_time"`

	Effects     []*effect.Effect                `json:"effects"`
	TypeGeoData []*obstacle_point.ObstaclePoint `json:"type_geo_data"`
	GeoData     []*obstacle_point.ObstaclePoint `json:"geo_data"`
}

func (o *Object) CalculateScale() {
	// подравниваем хп под размер туши
	percentLife := (o.HP / o.MaxHP) * 100                                         // сохраняем хп в процентном соотношение
	o.MaxHP = int(float64(o.TypeMaxHP) * float64(float64(o.Scale)/float64(100)))  // смотрим размер обьекта от оригинала и высчитываем его макс хп
	o.HP = int(float64(percentLife) * float64((float64(o.Scale) / float64(100)))) // востанавливаем хп в % соотношение

	// подравниваем тени под тушу
	//println(o.TypeXShadowOffset, o.TypeYShadowOffset)
	o.XShadowOffset = int(float64(o.TypeXShadowOffset) * (float64(o.Scale) / 100))
	o.YShadowOffset = int(float64(o.TypeYShadowOffset) * (float64(o.Scale) / 100))
	//println(o.XShadowOffset, o.YShadowOffset)

	o.SetGeoData()
}

func (o *Object) SetGeoData() {

	err := deepcopy.Copy(&o.GeoData, &o.TypeGeoData)
	if err != nil {
		println(err.Error())
	}

	for _, geoPoint := range o.GeoData {
		// применяем размер обьекта к геодате
		geoPoint.Radius = int(float64(geoPoint.Radius) * (float64(o.Scale) / 100))
		geoPoint.X = int(float64(geoPoint.X) * (float64(o.Scale) / 100))
		geoPoint.Y = int(float64(geoPoint.Y) * (float64(o.Scale) / 100))

		// получаем позицию гео точки на карте
		geoPoint.X += o.X
		geoPoint.Y += o.Y

		// поворачиваем геодату на угол обьекта
		newX, newY := game_math.RotatePoint(float64(geoPoint.X), float64(geoPoint.Y), float64(o.X),
			float64(o.Y), o.Rotate)

		geoPoint.X = int(newX)
		geoPoint.Y = int(newY)
	}
}

type Flore struct {
	TextureOverFlore string `json:"texture_over_flore"`
	TexturePriority  int    `json:"texture_priority"`
	X                int    `json:"x"`
	Y                int    `json:"y"`
}
