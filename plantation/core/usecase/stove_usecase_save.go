package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type StoveUseCaseSave struct {
	repository port.StoveIRepository
}

func NewStoveUseCaseSave(repository port.StoveIRepository) *StoveUseCaseSave {
	return &StoveUseCaseSave{repository: repository}
}

func (o *StoveUseCaseSave) Execute(Stove *entity.Stove) (*entity.Stove, error) {

	if err := Stove.IsValid(); err != nil {
		return nil, err
	}

	if i, err := o.repository.Save(Stove); err == nil && i != nil {
		return i.(*entity.Stove), err
	} else {
		return nil, err
	}
}
