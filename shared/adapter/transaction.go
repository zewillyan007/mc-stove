package adapter

import (
	"mc-stove/shared/connection/audit"
	"mc-stove/shared/port"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Db       *gorm.DB
	Envelope *audit.Envelope
}

func BeginTrans(provider port.IResourceProvider, typ audit.EventType) *Transaction {
	return BeginTransaction(provider.GetDB(), provider.RootRouteStr(), typ)
}

func BeginTransaction(db *gorm.DB, name string, typ audit.EventType) *Transaction {
	tx := &Transaction{Db: db.Begin()}

	if audit.DefaultPostman != nil {
		tx.Envelope = audit.DefaultPostman.NewEnvelope(name, typ)
	}
	return tx
}

func (o *Transaction) Commit() (db *gorm.DB, err error) {
	db = o.Db.Commit()
	if audit.DefaultPostman != nil {
		o.Envelope.SetEndTime(time.Now())
		err = audit.DefaultPostman.Push(o.Envelope)
	}
	return db, err
}

func (o *Transaction) Rollback(txError error) (db *gorm.DB, err error) {
	db = o.Db.Rollback()
	if audit.DefaultPostman != nil {
		o.Envelope.SetEndTime(time.Now())
		if txError != nil {
			o.Envelope.Options().Error = txError.Error()
		}
		err = audit.DefaultPostman.Push(o.Envelope)
	}
	return db, err
}

func (o *Transaction) GetTransaction() *gorm.DB {
	return o.Db
}

func (o *Transaction) GetEnvelope() *audit.Envelope {
	return o.Envelope
}
