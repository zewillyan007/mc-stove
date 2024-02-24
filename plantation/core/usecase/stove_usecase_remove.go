package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type StoveUseCaseRemove struct {
	repository port.StoveIRepository
}

func NewStoveUseCaseRemove(repository port.StoveIRepository) *StoveUseCaseRemove {
	return &StoveUseCaseRemove{repository: repository}
}

func (o *StoveUseCaseRemove) Execute(Stove *entity.Stove) error {
	return o.repository.Remove(Stove)
}
