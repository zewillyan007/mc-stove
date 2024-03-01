package service

import "mc-stove/plantation/core/domain/entity"

func FactoryStove() *entity.Stove {
	return entity.NewStove()
}

func FactoryPlant() *entity.Plant {
	return entity.NewPlant()
}

func FactoryTypeMicrocontroller() *entity.TypeMicrocontroller {
	return entity.NewTypeMicrocontroller()
}
