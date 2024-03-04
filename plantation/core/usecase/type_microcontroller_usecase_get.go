package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type TypeMicrocontrollerUseCaseGet struct {
	repository port.TypeMicrocontrollerIRepository
}

func NewTypeMicrocontrollerUseCaseGet(repository port.TypeMicrocontrollerIRepository) *TypeMicrocontrollerUseCaseGet {
	return &TypeMicrocontrollerUseCaseGet{repository: repository}
}

func (o *TypeMicrocontrollerUseCaseGet) Execute(id int64) (*entity.TypeMicrocontroller, error) {

	if i, err := o.repository.Get(id); err == nil && i != nil {
		return i.(*entity.TypeMicrocontroller), nil
	} else {
		return nil, err
	}
}
