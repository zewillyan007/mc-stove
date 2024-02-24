package service

import "mc-stove/plantation/core/domain/entity"

func FactoryStove() *entity.Stove {
	return entity.NewStove()
}
