package gameTypes

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/get"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"math/rand"
)

type resourceStore struct {
	base         map[int]resource.Resource
	recycled     map[int]resource.RecycledResource
	detail       map[int]resource.CraftDetail
	mapReservoir map[int]resource.Map
}

var Resource = newResourceStore()

func newResourceStore() *resourceStore {
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

func (r *resourceStore) GetAllDetails() map[int]resource.CraftDetail {
	return r.detail
}

func (r *resourceStore) GetAllBaseResource() map[int]resource.Resource {
	return r.base
}

func (r *resourceStore) GetAllRecycled() map[int]resource.RecycledResource {
	return r.recycled
}

func (r *resourceStore) GetRecycledByID(id int) (*resource.RecycledResource, bool) {
	var newResource resource.RecycledResource
	newResource, ok := r.recycled[id]
	return &newResource, ok
}

func (r *resourceStore) GetRecycledByName(name string) *resource.RecycledResource {
	for _, recycleRes := range r.recycled {
		if recycleRes.Name == name {
			return &recycleRes
		}
	}

	return nil
}

func (r *resourceStore) GetDetailByID(id int) (*resource.CraftDetail, bool) {
	var newResource resource.CraftDetail
	newResource, ok := r.detail[id]
	return &newResource, ok
}

func (r *resourceStore) GetDetailByName(name string) *resource.CraftDetail {
	for _, detail := range r.detail {
		if detail.Name == name {
			return &detail
		}
	}

	return nil
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

func (r *resourceStore) GetRandomMapResource() *resource.Map {
	allTypeResource := make([]resource.Map, 0)

	for _, typeRes := range r.mapReservoir {
		allTypeResource = append(allTypeResource, typeRes)
	}

	randomIndex := rand.Intn(len(allTypeResource))
	newResourceMap, _ := r.GetMapReservoirByID(allTypeResource[randomIndex].TypeID)

	return newResourceMap
}
