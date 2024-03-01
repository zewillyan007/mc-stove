package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type TypeMicrocontrollerUseCaseSave struct {
	repository port.TypeMicrocontrollerIRepository
}

func NewTypeMicrocontrollerUseCaseSave(repository port.TypeMicrocontrollerIRepository) *TypeMicrocontrollerUseCaseSave {
	return &TypeMicrocontrollerUseCaseSave{repository: repository}
}

func (o *TypeMicrocontrollerUseCaseSave) Execute(TypeMicrocontroller *entity.TypeMicrocontroller) (*entity.TypeMicrocontroller, error) {

	if err := TypeMicrocontroller.IsValid(); err != nil {
		return nil, err
	}

	if i, err := o.repository.Save(TypeMicrocontroller); err == nil && i != nil {
		return i.(*entity.TypeMicrocontroller), err
	} else {
		return nil, err
	}
}
