package port

import (
	"mc-stove/shared/types"

	"gorm.io/gorm"
)

type IResourceProvider interface {
	Context() *types.ManagerContext
	GetDB() *gorm.DB
	// GetCache() ICache
	RootRouteStr() string
}
