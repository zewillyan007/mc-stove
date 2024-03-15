package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type StovePlantUseCaseRemove struct {
	repository port.StovePlantIRepository
}

func NewStovePlantUseCaseRemove(repository port.StovePlantIRepository) *StovePlantUseCaseRemove {
	return &StovePlantUseCaseRemove{repository: repository}
}

func (o *StovePlantUseCaseRemove) Execute(StovePlant *entity.StovePlant) error {
	return o.repository.Remove(StovePlant)
}
