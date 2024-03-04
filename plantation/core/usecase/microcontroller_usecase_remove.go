package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type MicrocontrollerUseCaseRemove struct {
	repository port.MicrocontrollerIRepository
}

func NewMicrocontrollerUseCaseRemove(repository port.MicrocontrollerIRepository) *MicrocontrollerUseCaseRemove {
	return &MicrocontrollerUseCaseRemove{repository: repository}
}

func (o *MicrocontrollerUseCaseRemove) Execute(Microcontroller *entity.Microcontroller) error {
	return o.repository.Remove(Microcontroller)
}
