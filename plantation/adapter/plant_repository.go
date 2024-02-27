package adapter

import (
	"mc-stove/plantation/core/domain/entity"
	port2 "mc-stove/plantation/core/port"
	"mc-stove/shared/adapter"
	"mc-stove/shared/port"
	"mc-stove/shared/types"

	"gorm.io/gorm"
)

type PlantRepository struct {
	adapter.RepositoryCRUD
}

func NewPlantRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &PlantRepository{}
	repo.EntityType = (*entity.Plant)(nil)
	repo.SetTable("plantation.plant")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.PlantIRepository)(nil), func(args ...interface{}) interface{} {
		return NewPlantRepository(args[0].(*gorm.DB))
	})
}
