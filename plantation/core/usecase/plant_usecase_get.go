package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type PlantUseCaseGet struct {
	repository port.PlantIRepository
}

func NewPlantUseCaseGet(repository port.PlantIRepository) *PlantUseCaseGet {
	return &PlantUseCaseGet{repository: repository}
}

func (o *PlantUseCaseGet) Execute(id int64) (*entity.Plant, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.Plant), nil
	} else {
		return nil, err
	}
}
