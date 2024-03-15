package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type CurrentParamsUseCaseRemove struct {
	repository port.CurrentParamsIRepository
}

func NewCurrentParamsUseCaseRemove(repository port.CurrentParamsIRepository) *CurrentParamsUseCaseRemove {
	return &CurrentParamsUseCaseRemove{repository: repository}
}

func (o *CurrentParamsUseCaseRemove) Execute(CurrentParams *entity.CurrentParams) error {
	return o.repository.Remove(CurrentParams)
}
