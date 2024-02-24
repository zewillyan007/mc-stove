package adapter

import (
	"mc-stove/plantation/core/domain/entity"
	port2 "mc-stove/plantation/core/port"
	"mc-stove/shared/adapter"
	"mc-stove/shared/port"
	"mc-stove/shared/types"

	"gorm.io/gorm"
)

type StoveRepository struct {
	adapter.RepositoryCRUD
}

func NewStoveRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &StoveRepository{}
	repo.EntityType = (*entity.Stove)(nil)
	repo.SetTable("plantation.stove")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.StoveIRepository)(nil), func(args ...interface{}) interface{} {
		return NewStoveRepository(args[0].(*gorm.DB))
	})
}
