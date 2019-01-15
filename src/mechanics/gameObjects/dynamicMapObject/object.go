package dynamicMapObject

import (
	"../dialog"
	"time"
)

type DynamicObject struct {
	Dialog        *dialog.Dialog `json:"dialog"`
	TextureObject string         `json:"texture_object"`
	Destroyed     bool           `json:"destroyed"`
	Shadow        int            `json:"shadow"`
	DestroyTime   time.Time      `json:"destroy_time"`
}
