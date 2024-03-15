package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type MicrocontrollerStoveUseCaseGetAll struct {
	repository port.MicrocontrollerStoveIRepository
}

func NewMicrocontrollerStoveUseCaseGetAll(repository port.MicrocontrollerStoveIRepository) *MicrocontrollerStoveUseCaseGetAll {
	return &MicrocontrollerStoveUseCaseGetAll{repository: repository}
}

func (o *MicrocontrollerStoveUseCaseGetAll) Execute(conditions ...interface{}) []*entity.MicrocontrollerStove {

	list := o.repository.GetAll(conditions...)
	entities := make([]*entity.MicrocontrollerStove, len(list))
	for n, i := range list {
		entities[n] = i.(*entity.MicrocontrollerStove)
	}
	return entities
}
