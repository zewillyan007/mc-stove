package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type StovePlantUseCaseGetAll struct {
	repository port.StovePlantIRepository
}

func NewStovePlantUseCaseGetAll(repository port.StovePlantIRepository) *StovePlantUseCaseGetAll {
	return &StovePlantUseCaseGetAll{repository: repository}
}

func (o *StovePlantUseCaseGetAll) Execute(conditions ...interface{}) []*entity.StovePlant {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.StovePlant, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.StovePlant)
	}
	return entities
}
