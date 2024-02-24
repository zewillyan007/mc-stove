package resource

import (
	"mc-stove/shared/types"

	"gorm.io/gorm"
)

type ResourceProvider struct {
	db   *gorm.DB
	ctxt *types.ManagerContext
	// cache        port.ICache
	RootRoute    string
	CurrentRoute string
}

// com cache
//
//	func NewResourceProvider(db *gorm.DB, ctxt *types.ManagerContext, cache port.ICache) *ResourceProvider {
//		return &ResourceProvider{
//			db:           db,
//			ctxt:         ctxt,
//			cache:        cache,
//			RootRoute:    "",
//			CurrentRoute: "",
//		}
//	}
func NewResourceProvider(db *gorm.DB, ctxt *types.ManagerContext) *ResourceProvider {
	return &ResourceProvider{
		db:           db,
		ctxt:         ctxt,
		RootRoute:    "",
		CurrentRoute: "",
	}
}

func (r *ResourceProvider) Context() *types.ManagerContext {
	return r.ctxt
}

func (r *ResourceProvider) GetDB() *gorm.DB {
	return r.db
}

// func (r *ResourceProvider) GetCache() port.ICache {
// 	return r.cache
// }

func (r *ResourceProvider) RootRouteStr() string {
	return r.RootRoute
}
