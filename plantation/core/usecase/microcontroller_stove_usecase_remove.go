package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type MicrocontrollerStoveUseCaseRemove struct {
	repository port.MicrocontrollerStoveIRepository
}

func NewMicrocontrollerStoveUseCaseRemove(repository port.MicrocontrollerStoveIRepository) *MicrocontrollerStoveUseCaseRemove {
	return &MicrocontrollerStoveUseCaseRemove{repository: repository}
}

func (o *MicrocontrollerStoveUseCaseRemove) Execute(MicrocontrollerStove *entity.MicrocontrollerStove) error {
	return o.repository.Remove(MicrocontrollerStove)
}
