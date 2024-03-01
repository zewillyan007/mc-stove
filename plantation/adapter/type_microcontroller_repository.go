package adapter

import (
	"mc-stove/plantation/core/domain/entity"
	port2 "mc-stove/plantation/core/port"
	"mc-stove/shared/adapter"
	"mc-stove/shared/port"
	"mc-stove/shared/types"

	"gorm.io/gorm"
)

type TypeMicrocontrollerRepository struct {
	adapter.RepositoryCRUD
}

func NewTypeMicrocontrollerRepository(db *gorm.DB) port.IRepositoryCRUD {
	repo := &TypeMicrocontrollerRepository{}
	repo.EntityType = (*entity.TypeMicrocontroller)(nil)
	repo.SetTable("domain.type_microcontroller")
	repo.SetDB(db)

	return repo
}

func init() {
	types.SetConstructor((*port2.TypeMicrocontrollerIRepository)(nil), func(args ...interface{}) interface{} {
		return NewTypeMicrocontrollerRepository(args[0].(*gorm.DB))
	})
}
