package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type StoveUseCaseGetAll struct {
	repository port.StoveIRepository
}

func NewStoveUseCaseGetAll(repository port.StoveIRepository) *StoveUseCaseGetAll {
	return &StoveUseCaseGetAll{repository: repository}
}

func (o *StoveUseCaseGetAll) Execute(conditions ...interface{}) []*entity.Stove {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.Stove, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.Stove)
	}
	return entities
}
