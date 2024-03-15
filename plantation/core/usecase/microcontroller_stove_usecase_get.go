package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type MicrocontrollerStoveUseCaseGet struct {
	repository port.MicrocontrollerStoveIRepository
}

func NewMicrocontrollerStoveUseCaseGet(repository port.MicrocontrollerStoveIRepository) *MicrocontrollerStoveUseCaseGet {
	return &MicrocontrollerStoveUseCaseGet{repository: repository}
}

func (o *MicrocontrollerStoveUseCaseGet) Execute(id int64) (*entity.MicrocontrollerStove, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.MicrocontrollerStove), nil
	} else {
		return nil, err
	}
}
