package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type StovePlantUseCaseGet struct {
	repository port.StovePlantIRepository
}

func NewStovePlantUseCaseGet(repository port.StovePlantIRepository) *StovePlantUseCaseGet {
	return &StovePlantUseCaseGet{repository: repository}
}

func (o *StovePlantUseCaseGet) Execute(id int64) (*entity.StovePlant, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.StovePlant), nil
	} else {
		return nil, err
	}
}
