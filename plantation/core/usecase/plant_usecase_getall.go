package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type PlantUseCaseGetAll struct {
	repository port.PlantIRepository
}

func NewPlantUseCaseGetAll(repository port.PlantIRepository) *PlantUseCaseGetAll {
	return &PlantUseCaseGetAll{repository: repository}
}

func (o *PlantUseCaseGetAll) Execute(conditions ...interface{}) []*entity.Plant {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.Plant, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.Plant)
	}
	return entities
}
