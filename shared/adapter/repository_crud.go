package adapter

import (
	"database/sql"
	"mc-stove/shared/port"
	"reflect"

	"gorm.io/gorm"
)

type RepositoryCRUD struct {
	Repository
	EntityType port.IEntity
}

func NewRepositoryCRUD(db *gorm.DB) port.IRepositoryCRUD {
	repo := &RepositoryCRUD{}
	repo.db = db
	return repo
}

func (crud *RepositoryCRUD) Save(entity port.IEntity) (port.IEntity, error) {
	if entity.GetId() == 0 {
		return entity, crud.Insert(entity)
	} else {
		return entity, crud.Update(entity)
	}
}

func (crud *RepositoryCRUD) Get(id int64) (port.IEntity, error) {
	ent := crud.NewEntity()
	crud.First(ent, id)
	return ent, nil
}

func (crud *RepositoryCRUD) GetAll(conditions ...interface{}) []port.IEntity {

	list := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(crud.EntityType)), 0, 0).Interface()
	crud.Find(&list, conditions...)

	reflectValue := reflect.ValueOf(list)
	entities := make([]port.IEntity, reflectValue.Len())
	for n := 0; n < len(entities); n++ {
		i := reflectValue.Index(n).Addr().Elem().Interface()
		entities[n] = i.(port.IEntity)
	}

	return entities
}

func (crud *RepositoryCRUD) Remove(entity port.IEntity) error {
	return crud.Repository.Delete(entity)
}

func (crud *RepositoryCRUD) NewEntity() port.IEntity {
	return reflect.New(reflect.TypeOf(crud.EntityType).Elem()).Interface().(port.IEntity)
}

func (crud *RepositoryCRUD) SetEntityType(typ port.IEntity) {
	crud.EntityType = typ
}

func (crud *RepositoryCRUD) SqlQueryRow(sql string) *sql.Row {
	return crud.QueryRow(sql)
}

func (crud *RepositoryCRUD) SqlQueryRows(sql string) (*sql.Rows, error) {
	return crud.QueryRows(sql)
}

func (crud *RepositoryCRUD) SqlQueryPaginator(columns string, table string, where string, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	return crud.QueryPaginator(columns, table, where, sqlTemplate, page, limit, order...)
}

func (crud *RepositoryCRUD) WithTransaction(transaction port.ITransaction) port.IRepositoryCRUD {
	crud.SetDB(transaction.GetTransaction())
	// crud.Envelope = transaction.GetEnvelope()
	return crud
}

//func (crud *RepositoryCRUD) WithTransaction(transaction port.ITransaction) port.IRepositoryCRUD {
//	i := reflect.New(reflect.TypeOf(crud).Elem()).Interface()
//	repo := i.(port.IRepositoryCRUD)
//	repo.SetEntityType(crud.EntityType)
//	repo.SetTable(crud.GetTable())
//	repo.SetDB(transaction.GetTransaction())
//
//	return repo
//}

func (crud *RepositoryCRUD) SqlQueryData(columns, table, where, sqlTemplate string, order ...string) (map[string]interface{}, error) {
	return crud.QueryData(columns, table, where, sqlTemplate, order...)
}
