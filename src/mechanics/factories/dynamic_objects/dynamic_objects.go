package dynamic_objects

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dynamic_map_object"
	"github.com/getlantern/deepcopy"
	"sync"
)

var DynamicObjects = newObjectsStore()

type store struct {
	mx      sync.Mutex
	objects map[int]*dynamic_map_object.Object
}

func newObjectsStore() *store {
	object := maps.AllTypeCoordinate()
	mapObjects := make(map[int]*dynamic_map_object.Object)

	for _, obj := range object {
		mapObjects[obj.TypeID] = obj
	}

	return &store{
		objects: mapObjects,
	}
}

func (d *store) GetDynamicObjectByID(id int, rotate int) *dynamic_map_object.Object {
	d.mx.Lock()
	defer d.mx.Unlock()

	var newObj dynamic_map_object.Object
	factoryObj, ok := d.objects[id]

	if ok {
		err := deepcopy.Copy(&newObj, &factoryObj)
		if err != nil {
			println(err.Error())
		}

		newObj.Rotate = rotate
		return &newObj
	} else {
		return nil
	}
}

func (d *store) GetDynamicObjectByTexture(name string, rotate int) *dynamic_map_object.Object {
	d.mx.Lock()
	defer d.mx.Unlock()

	var newObj dynamic_map_object.Object

	for _, factoryObj := range d.objects {
		if factoryObj.Texture == name {
			err := deepcopy.Copy(&newObj, &factoryObj)
			if err != nil {
				println(err.Error())
			}

			newObj.Rotate = rotate
			return &newObj
		}
	}

	return nil
}
