package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type TypeMicrocontrollerUseCaseGetAll struct {
	repository port.TypeMicrocontrollerIRepository
}

func NewTypeMicrocontrollerUseCaseGetAll(repository port.TypeMicrocontrollerIRepository) *TypeMicrocontrollerUseCaseGetAll {
	return &TypeMicrocontrollerUseCaseGetAll{repository: repository}
}

func (o *TypeMicrocontrollerUseCaseGetAll) Execute(conditions ...interface{}) []*entity.TypeMicrocontroller {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.TypeMicrocontroller, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.TypeMicrocontroller)
	}
	return entities
}
