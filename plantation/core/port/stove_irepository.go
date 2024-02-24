package port

import (
	port_shared "mc-stove/shared/port"
)

type StoveIRepository interface {
	port_shared.IRepositoryCRUD
	//GetAll(conditions ...interface{}) []*entity.Company
	//Remove(*entity.Company) error
	//Get(int64) (*entity.Company, error)
	//Save(*entity.Company) (*entity.Company, error)
	//SqlQueryRow(string) *sql.Row
	//SqlQueryRows(string) (*sql.Rows, error)
	//SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error)
	//WithTransaction(transaction port_shared.ITransaction) CompanyIRepository
	//SqlQueryData(columns, table, where, sqlTemplate string, order ...string) (map[string]interface{}, error)
}
