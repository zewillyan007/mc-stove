package port

import (
	"mc-stove/shared/connection/audit"

	"gorm.io/gorm"
)

type ITransaction interface {
	GetTransaction() *gorm.DB
	GetEnvelope() *audit.Envelope
}
