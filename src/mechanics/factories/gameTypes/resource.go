package gameTypes

import (
	"../../db/get"
	"../../gameObjects/resource"
)

type resourceStore struct {
	base map[int]resource.Resource
	recycled map[int]resource.RecycledResource
}

var Resource = NewResourceStore()

func NewResourceStore() *resourceStore {
	return &resourceStore{base: get.ResourceType(), recycled: get.RecycledResourceType()}
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