package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type CurrentParamsUseCaseGetAll struct {
	repository port.CurrentParamsIRepository
}

func NewCurrentParamsUseCaseGetAll(repository port.CurrentParamsIRepository) *CurrentParamsUseCaseGetAll {
	return &CurrentParamsUseCaseGetAll{repository: repository}
}

func (o *CurrentParamsUseCaseGetAll) Execute(conditions ...interface{}) []*entity.CurrentParams {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.CurrentParams, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.CurrentParams)
	}
	return entities
}
