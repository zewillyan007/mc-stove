package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type StovePlantUseCaseSave struct {
	repository port.StovePlantIRepository
}

func NewStovePlantUseCaseSave(repository port.StovePlantIRepository) *StovePlantUseCaseSave {
	return &StovePlantUseCaseSave{repository: repository}
}

func (o *StovePlantUseCaseSave) Execute(StovePlant *entity.StovePlant) (*entity.StovePlant, error) {

	if err := StovePlant.IsValid(); err != nil {
		return nil, err
	}

	if i, err := o.repository.Save(StovePlant); err == nil && i != nil {
		return i.(*entity.StovePlant), err
	} else {
		return nil, err
	}
}
