package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type PlantUseCaseSave struct {
	repository port.PlantIRepository
}

func NewPlantUseCaseSave(repository port.PlantIRepository) *PlantUseCaseSave {
	return &PlantUseCaseSave{repository: repository}
}

func (o *PlantUseCaseSave) Execute(Plant *entity.Plant) (*entity.Plant, error) {

	if err := Plant.IsValid(); err != nil {
		return nil, err
	}

	if i, err := o.repository.Save(Plant); err == nil && i != nil {
		return i.(*entity.Plant), err
	} else {
		return nil, err
	}
}
