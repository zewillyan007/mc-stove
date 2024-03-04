package adapter

import (
	"mc-stove/plantation/core/domain/entity"
	port2 "mc-stove/plantation/core/port"
	"mc-stove/shared/adapter"
	"mc-stove/shared/port"
	"mc-stove/shared/types"

	"gorm.io/gorm"
)

type MicrocontrollerRepository struct {
	adapter.RepositoryCRUD
}

func NewMicrocontrollerRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &MicrocontrollerRepository{}
	repo.EntityType = (*entity.Microcontroller)(nil)
	repo.SetTable("plantation.microcontroller")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.MicrocontrollerIRepository)(nil), func(args ...interface{}) interface{} {
		return NewMicrocontrollerRepository(args[0].(*gorm.DB))
	})
}
