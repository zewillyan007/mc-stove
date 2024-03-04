package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type MicrocontrollerUseCaseGetAll struct {
	repository port.MicrocontrollerIRepository
}

func NewMicrocontrollerUseCaseGetAll(repository port.MicrocontrollerIRepository) *MicrocontrollerUseCaseGetAll {
	return &MicrocontrollerUseCaseGetAll{repository: repository}
}

func (o *MicrocontrollerUseCaseGetAll) Execute(conditions ...interface{}) []*entity.Microcontroller {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.Microcontroller, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.Microcontroller)
	}
	return entities
}
