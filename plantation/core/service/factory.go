package service

import "mc-stove/plantation/core/domain/entity"

func FactoryStove() *entity.Stove {
	return entity.NewStove()
}

func FactoryPlant() *entity.Plant {
	return entity.NewPlant()
}

func FactoryMicrocontroller() *entity.Microcontroller {
	return entity.NewMicrocontroller()
}

func FactoryTypeMicrocontroller() *entity.TypeMicrocontroller {
	return entity.NewTypeMicrocontroller()
}

func FactoryCurrentParams() *entity.CurrentParams {
	return entity.NewCurrentParams()
}

func FactoryStovePlant() *entity.StovePlant {
	return entity.NewStovePlant()
}

func FactoryMicrocontrollerStove() *entity.MicrocontrollerStove {
	return entity.NewMicrocontrollerStove()
}
