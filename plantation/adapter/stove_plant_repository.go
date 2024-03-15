package adapter

import (
	"mc-stove/plantation/core/domain/entity"
	port2 "mc-stove/plantation/core/port"
	"mc-stove/shared/adapter"
	"mc-stove/shared/port"
	"mc-stove/shared/types"

	"gorm.io/gorm"
)

type StovePlantRepository struct {
	adapter.RepositoryCRUD
}

func NewStovePlantRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &StovePlantRepository{}
	repo.EntityType = (*entity.StovePlant)(nil)
	repo.SetTable("plantation.stove_plant")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.StovePlantIRepository)(nil), func(args ...interface{}) interface{} {
		return NewStovePlantRepository(args[0].(*gorm.DB))
	})
}
