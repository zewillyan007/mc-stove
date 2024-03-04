package usecase

import (
	"mc-stove/plantation/core/domain/entity"
	"mc-stove/plantation/core/port"
)

type TypeMicrocontrollerUseCaseRemove struct {
	repository port.TypeMicrocontrollerIRepository
}

func NewTypeMicrocontrollerUseCaseRemove(repository port.TypeMicrocontrollerIRepository) *TypeMicrocontrollerUseCaseRemove {
	return &TypeMicrocontrollerUseCaseRemove{repository: repository}
}

func (o *TypeMicrocontrollerUseCaseRemove) Execute(TypeMicrocontroller *entity.TypeMicrocontroller) error {
	return o.repository.Remove(TypeMicrocontroller)
}
