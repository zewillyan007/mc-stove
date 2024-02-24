package port

import "database/sql"

type IRepositoryCRUD interface {
	IRepository
	SetEntityType(typ IEntity)
	Save(entity IEntity) (IEntity, error)
	Get(id int64) (IEntity, error)
	GetAll(conditions ...interface{}) []IEntity
	Remove(entity IEntity) error
	SqlQueryRow(string) *sql.Row
	SqlQueryRows(string) (*sql.Rows, error)
	SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error)
	SqlQueryData(columns, table, where, sqlTemplate string, order ...string) (map[string]interface{}, error)
	WithTransaction(transaction ITransaction) IRepositoryCRUD
}
