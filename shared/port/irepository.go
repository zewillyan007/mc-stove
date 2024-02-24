package port

import "gorm.io/gorm"

type IRepository interface {
	SetTable(name string)
	SetDB(db *gorm.DB)
	SetContext(ctxt IManagerContext)
	GetContext() IManagerContext
}
