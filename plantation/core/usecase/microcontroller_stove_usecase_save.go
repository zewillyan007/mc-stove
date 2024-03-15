package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type MicrocontrollerStoveUseCaseSave struct {
	repository port.MicrocontrollerStoveIRepository
}

func NewMicrocontrollerStoveUseCaseSave(repository port.MicrocontrollerStoveIRepository) *MicrocontrollerStoveUseCaseSave {
	return &MicrocontrollerStoveUseCaseSave{repository: repository}
}

func (o *MicrocontrollerStoveUseCaseSave) Execute(MicrocontrollerStove *entity.MicrocontrollerStove) (*entity.MicrocontrollerStove, error) {

	if err := MicrocontrollerStove.IsValid(); err != nil {
		return nil, err
	}

	if i, err := o.repository.Save(MicrocontrollerStove); err == nil && i != nil {
		return i.(*entity.MicrocontrollerStove), err
	} else {
		return nil, err
	}
}
