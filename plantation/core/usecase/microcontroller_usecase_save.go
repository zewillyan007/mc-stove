package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type MicrocontrollerUseCaseSave struct {
	repository port.MicrocontrollerIRepository
}

func NewMicrocontrollerUseCaseSave(repository port.MicrocontrollerIRepository) *MicrocontrollerUseCaseSave {
	return &MicrocontrollerUseCaseSave{repository: repository}
}

func (o *MicrocontrollerUseCaseSave) Execute(Microcontroller *entity.Microcontroller) (*entity.Microcontroller, error) {

	if err := Microcontroller.IsValid(); err != nil {
		return nil, err
	}

	if i, err := o.repository.Save(Microcontroller); err == nil && i != nil {
		return i.(*entity.Microcontroller), err
	} else {
		return nil, err
	}
}
