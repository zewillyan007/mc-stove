package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type CurrentParamsUseCaseSave struct {
	repository port.CurrentParamsIRepository
}

func NewCurrentParamsUseCaseSave(repository port.CurrentParamsIRepository) *CurrentParamsUseCaseSave {
	return &CurrentParamsUseCaseSave{repository: repository}
}

func (o *CurrentParamsUseCaseSave) Execute(CurrentParams *entity.CurrentParams) (*entity.CurrentParams, error) {

	if err := CurrentParams.IsValid(); err != nil {
		return nil, err
	}

	if i, err := o.repository.Save(CurrentParams); err == nil && i != nil {
		return i.(*entity.CurrentParams), err
	} else {
		return nil, err
	}
}
