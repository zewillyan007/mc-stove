package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type StoveUseCaseGet struct {
	repository port.StoveIRepository
}

func NewStoveUseCaseGet(repository port.StoveIRepository) *StoveUseCaseGet {
	return &StoveUseCaseGet{repository: repository}
}

func (o *StoveUseCaseGet) Execute(id int64) (*entity.Stove, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.Stove), nil
	} else {
		return nil, err
	}
}
