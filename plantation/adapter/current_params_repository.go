package adapter

import (
	"mc-stove/plantation/core/domain/entity"
	port2 "mc-stove/plantation/core/port"
	"mc-stove/shared/adapter"
	"mc-stove/shared/port"
	"mc-stove/shared/types"

	"gorm.io/gorm"
)

type CurrentParamsRepository struct {
	adapter.RepositoryCRUD
}

func NewCurrentParamsRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &CurrentParamsRepository{}
	repo.EntityType = (*entity.CurrentParams)(nil)
	repo.SetTable("plantation.stove_current_params")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.CurrentParamsIRepository)(nil), func(args ...interface{}) interface{} {
		return NewCurrentParamsRepository(args[0].(*gorm.DB))
	})
}
