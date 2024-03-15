package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type CurrentParamsUseCaseGet struct {
	repository port.CurrentParamsIRepository
}

func NewCurrentParamsUseCaseGet(repository port.CurrentParamsIRepository) *CurrentParamsUseCaseGet {
	return &CurrentParamsUseCaseGet{repository: repository}
}

func (o *CurrentParamsUseCaseGet) Execute(id int64) (*entity.CurrentParams, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.CurrentParams), nil
	} else {
		return nil, err
	}
}
