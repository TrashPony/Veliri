package dynamicMapObject

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/effect"
	"time"
)

type DynamicObject struct {
	Dialog *dialog.Dialog `json:"dialog"`

	TextureBackground string `json:"texture_background"`
	BackgroundScale   int    `json:"background_scale"`
	BackgroundRotate  int    `json:"background_rotate"`

	TextureObject       string `json:"texture_object"`
	AnimateSpriteSheets string `json:"animate_sprite_sheets"`
	ObjectScale         int    `json:"object_scale"`
	ObjectRotate        int    `json:"object_rotate"`

	Move   bool `json:"move"`
	View   bool `json:"view"`
	Attack bool `json:"attack"`

	AnimateLoop bool      `json:"animate_loop"`
	Destroyed   bool      `json:"destroyed"`
	Shadow      int       `json:"shadow"`
	DestroyTime time.Time `json:"destroy_time"`

	Effects []*effect.Effect `json:"effects"`
}
