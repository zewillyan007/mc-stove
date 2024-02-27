package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type PlantUseCaseRemove struct {
	repository port.PlantIRepository
}

func NewPlantUseCaseRemove(repository port.PlantIRepository) *PlantUseCaseRemove {
	return &PlantUseCaseRemove{repository: repository}
}

func (o *PlantUseCaseRemove) Execute(Plant *entity.Plant) error {
	return o.repository.Remove(Plant)
}
