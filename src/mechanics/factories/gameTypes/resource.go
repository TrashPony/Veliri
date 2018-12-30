package gameTypes

import (
	"../../db/get"
	"../../gameObjects/resource"
	"math/rand"
)

type resourceStore struct {
	base         map[int]resource.Resource
	recycled     map[int]resource.RecycledResource
	detail       map[int]resource.CraftDetail
	mapReservoir map[int]resource.Map
}

var Resource = NewResourceStore()

func NewResourceStore() *resourceStore {
	return &resourceStore{
		base:         get.ResourceType(),
		recycled:     get.RecycledResourceType(),
		detail:       get.DetailType(),
		mapReservoir: get.ReservoirMapType(),
	}
}

func (r *resourceStore) GetBaseByID(id int) (*resource.Resource, bool) {
	var newResource resource.Resource
	newResource, ok := r.base[id]
	return &newResource, ok
}

func (r *resourceStore) GetRecycledByID(id int) (*resource.RecycledResource, bool) {
	var newResource resource.RecycledResource
	newResource, ok := r.recycled[id]
	return &newResource, ok
}

func (r *resourceStore) GetDetailByID(id int) (*resource.CraftDetail, bool) {
	var newResource resource.CraftDetail
	newResource, ok := r.detail[id]
	return &newResource, ok
}

func (r *resourceStore) GetMapReservoirByID(id int) (*resource.Map, bool) {
	newReservoir, ok := r.mapReservoir[id]

	baseRes, _ := r.GetBaseByID(newReservoir.ResourceID)
	newReservoir.Resource = baseRes

	newReservoir.Count = rand.Intn(newReservoir.MaxCount-newReservoir.MinCount) + newReservoir.MinCount

	return &newReservoir, ok
}

func (r *resourceStore) GetAllTypeMapResource() map[int]resource.Map {
	return r.mapReservoir
}
