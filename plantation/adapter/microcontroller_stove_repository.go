package adapter

import (
	"mc-stove/plantation/core/domain/entity"
	port2 "mc-stove/plantation/core/port"
	"mc-stove/shared/adapter"
	"mc-stove/shared/port"
	"mc-stove/shared/types"

	"gorm.io/gorm"
)

type MicrocontrollerStoveRepository struct {
	adapter.RepositoryCRUD
}

func NewMicrocontrollerStoveRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &MicrocontrollerStoveRepository{}
	repo.EntityType = (*entity.MicrocontrollerStove)(nil)
	repo.SetTable("plantation.microcontroller_stove")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.MicrocontrollerStoveIRepository)(nil), func(args ...interface{}) interface{} {
		return NewMicrocontrollerStoveRepository(args[0].(*gorm.DB))
	})
}
