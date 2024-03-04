package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type MicrocontrollerUseCaseGet struct {
	repository port.MicrocontrollerIRepository
}

func NewMicrocontrollerUseCaseGet(repository port.MicrocontrollerIRepository) *MicrocontrollerUseCaseGet {
	return &MicrocontrollerUseCaseGet{repository: repository}
}

func (o *MicrocontrollerUseCaseGet) Execute(id int64) (*entity.Microcontroller, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.Microcontroller), nil
	} else {
		return nil, err
	}
}
