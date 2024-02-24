package adapter

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"mc-stove/shared/connection/audit"
	"mc-stove/shared/constant"
	"mc-stove/shared/grid"
	"mc-stove/shared/port"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

type Repository struct {
	db       *gorm.DB
	table    string
	Envelope *audit.Envelope
	Context  port.IManagerContext
}

func NewRepository(db *gorm.DB, table string) *Repository {
	return &Repository{db: db, table: table}
}

func NewRepositoryWithTransaction(transaction port.ITransaction, table string) *Repository {
	return &Repository{db: transaction.GetTransaction(), table: table, Envelope: transaction.GetEnvelope()}
}

func (o *Repository) GetTable() string {
	return o.table
}

func (o *Repository) Db() *gorm.DB {
	return o.db
}

func (o *Repository) SetTable(name string) {
	o.table = name
}

func (o *Repository) SetDB(db *gorm.DB) {
	o.db = db
}

func (o *Repository) SetContext(ctxt port.IManagerContext) {
	o.Context = ctxt
}

func (o *Repository) GetContext() port.IManagerContext {
	return o.Context
}

func (o *Repository) Insert(value interface{}) error {
	return o.auditOperation(audit.Insert, value)
}

func (o *Repository) Update(value interface{}) error {
	return o.auditOperation(audit.Update, value)
}

func (o *Repository) Delete(value interface{}) error {
	return o.auditOperation(audit.Delete, value)
}

type fnExecutor = func(db *gorm.DB) *gorm.DB

func (o *Repository) auditOperation(operation audit.EventType, value interface{}) (err error) {

	var db *gorm.DB
	var stmt *gorm.Statement
	var result *gorm.DB
	var executor fnExecutor

	if audit.DefaultPostman == nil {

		switch operation {
		case audit.Insert:
			result = o.db.Table(o.table).Create(value)
		case audit.Update:
			result = o.db.Table(o.table).Save(value)
		case audit.Delete:
			result = o.db.Table(o.table).Delete(value)
		}

	} else {
		var envelope *audit.Envelope

		if o.isToSendEnvelope() {
			envelope = audit.DefaultPostman.NewEnvelope(o.table, operation)
		} else {
			envelope = o.Envelope
		}

		if o.Context != nil && o.Context.GetUser() != nil {
			envelope.Options().UserId = o.Context.GetUser().Id
			envelope.Options().UserName = o.Context.GetUser().Name
		}

		db = o.db.Session(&gorm.Session{DryRun: true})

		switch operation {
		case audit.Insert:
			stmt = db.Table(o.table).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).Create(value).Statement
			executor = stmt.Callback().Create().Execute
		case audit.Update:
			stmt = db.Table(o.table).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).Save(value).Statement
			executor = stmt.Callback().Update().Execute
		case audit.Delete:
			stmt = db.Table(o.table).Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).Delete(value).Statement
			executor = stmt.Callback().Delete().Execute
		}

		sqlText := stmt.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)

		envelope.AddTask(operation, o.table, 0, sqlText, value)
		if o.isToSendEnvelope() {
			if err = audit.DefaultPostman.Push(envelope); err != nil {
				return err
			}
		}

		stmt.Config.DryRun = false
		result = executor(stmt.DB)

		if o.isToSendEnvelope() {
			envelope.SetEndTime(time.Now())
			if result.Error != nil {
				envelope.Options().Error = result.Error.Error()
				envelope.SetStatus(audit.Error)
			} else {
				envelope.SetStatus(audit.Success)
			}
			if err = audit.DefaultPostman.Push(envelope); err != nil {
				return errors.New(envelope.Options().Error + "\n" + err.Error())
			}
		}

	}

	return result.Error
}

func (o *Repository) isToSendEnvelope() bool {
	return o.Envelope == nil
}

func (o *Repository) First(dest interface{}, conds ...interface{}) {
	if conds != nil {
		o.db.Table(o.table).Find(dest, conds)
	} else {
		o.db.Table(o.table).Find(dest)
	}
}

func (o *Repository) Find(dest interface{}, conds ...interface{}) {
	o.db.Table(o.table).Find(dest, conds...)
}

func (o *Repository) Order(cols interface{}) *gorm.DB {
	return o.db.Table(o.table).Order(cols)
}

func (o *Repository) Exec(sql string, values ...interface{}) (err error) {
	if audit.DefaultPostman == nil {
		return o.db.Exec(sql, values...).Error
	}

	sqlText := o.db.Statement.Dialector.Explain(sql, values...)
	var envelope *audit.Envelope
	typ := audit.DefaultPostman.ExtractEvent(sql)

	if o.isToSendEnvelope() {
		envelope = audit.DefaultPostman.NewEnvelope(o.table, typ)
	} else {
		envelope = o.Envelope
	}
	envelope.AddTask(typ, o.table, 0, sqlText, nil)
	if o.isToSendEnvelope() {
		if err = audit.DefaultPostman.Push(envelope); err != nil {
			return err
		}
	}

	err = o.db.Exec(sql, values...).Error

	if o.isToSendEnvelope() {
		envelope.SetEndTime(time.Now())
		if err != nil {
			envelope.Options().Error = err.Error()
			envelope.SetStatus(audit.Error)
		} else {
			envelope.SetStatus(audit.Success)
		}
		if er := audit.DefaultPostman.Push(envelope); er != nil {
			return errors.New(envelope.Options().Error + "\n" + er.Error())
		}
	}

	return err
}

func (o *Repository) QueryRow(sql string) *sql.Row {
	return o.db.Raw(sql).Row()
}

func (o *Repository) QueryRows(sql string) (*sql.Rows, error) {
	return o.db.Raw(sql).Rows()
}

func (o *Repository) QueryPaginatorGroupBy(columns, table, where, group, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	var oby string
	if len(order) > 0 {
		oby = order[0]
	}

	return o.QueryPaginatorWhenGroupBy(columns, table, where, sqlTemplate, page, limit, group, oby)
}

func (o *Repository) QueryPaginatorGroupByAndHaving(columns, table, where, group, having, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {
	var oby string
	if len(order) > 0 {
		oby = order[0]
	}

	return o.QueryPaginatorWhenGroupByAndHaving(columns, table, where, sqlTemplate, page, limit, group, having, oby)
}

func (o *Repository) QueryData(columns, table, where, sqlTemplate string, order ...string) (map[string]interface{}, error) {

	var sql string
	var orderBy string
	data := make(map[string]interface{})

	var sqlParsed string
	if len(where) > 0 {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table, where)
	} else {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table)
	}

	if len(order) > 0 && len(order[0]) > 0 {
		orderBy = " ORDER BY " + order[0]
	}

	sql = sqlParsed + orderBy + constant.MAX_RECORDS

	rows, err := o.QueryRows(sql)

	if err != nil {
		return nil, err
	}

	data["prev"] = 0
	data["next"] = 0
	data["page"] = 0
	data["pages"] = 0
	data["total"] = 0
	data["rows"] = grid.RowsToInterface(rows)
	data["lines"] = len(data["rows"].([]interface{}))

	return data, nil
}

func (o *Repository) QueryDataV2(columns, table, where, sqlTemplate string, limit float64, order ...string) (map[string]interface{}, error) {

	var sql string
	var orderBy string
	data := make(map[string]interface{})

	var sqlParsed string
	if len(where) > 0 {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table, where)
	} else {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table)
	}

	if len(order) > 0 && len(order[0]) > 0 {
		orderBy = " ORDER BY " + order[0]
	}

	sql = sqlParsed + orderBy + " OFFSET 0 LIMIT " + fmt.Sprintf("%v", limit)

	rows, err := o.QueryRows(sql)

	if err != nil {
		return nil, err
	}

	data["prev"] = 0
	data["next"] = 0
	data["page"] = 0
	data["pages"] = 0
	data["total"] = 0
	data["rows"] = grid.RowsToInterface(rows)
	data["lines"] = len(data["rows"].([]interface{}))

	return data, nil
}

func (o *Repository) QueryDataWhenGroupByAndHaving(columns, table, where, sqlTemplate string, limit float64, group, _having_, order string) (map[string]interface{}, error) {

	var stmt string
	var orderBy, groupBy, having string
	data := make(map[string]interface{})

	var sqlParsed string
	if len(where) > 0 {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table, where)
	} else {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table)
	}

	if len(group) > 0 {
		groupBy = " GROUP BY " + group
	}
	if len(_having_) > 0 {
		having = " HAVING" + _having_
	}
	if len(order) > 0 {
		orderBy = " ORDER BY " + order
	}

	if limit == 0 {
		return nil, fmt.Errorf("40020")
	}

	stmt = sqlParsed + groupBy + having + orderBy + " OFFSET 0 LIMIT " + fmt.Sprintf("%v", limit)
	rows, err := o.QueryRows(stmt)

	if err != nil {
		return nil, err
	}

	data["prev"] = 0
	data["next"] = 0
	data["page"] = 0
	data["pages"] = 0
	data["total"] = 0
	data["rows"] = grid.RowsToInterface(rows)
	data["lines"] = len(data["rows"].([]interface{}))

	return data, nil
}

func (o *Repository) QueryDataGroupByAndHaving(columns, table, where, group, having, sqlTemplate string, limit float64, order ...string) (map[string]interface{}, error) {
	var oby string
	if len(order) > 0 {
		oby = order[0]
	}

	return o.QueryDataWhenGroupByAndHaving(columns, table, where, sqlTemplate, limit, group, having, oby)
}

func (o *Repository) QueryPaginator(columns, table, where, sqlTemplate string, page float64, limit float64, order ...string) (map[string]interface{}, error) {

	var sql string
	var orderby string
	var total, offset, pages float64
	data := make(map[string]interface{})

	var sqlParsed string
	if len(where) > 0 {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table, where)
	} else {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table)
	}

	if limit > 0 {
		offset = (page - 1) * limit
		rowsCount := o.QueryRow("SELECT COUNT(*) FROM " + table + " WHERE " + where)
		rowsCount.Scan(&total)
		if len(order) > 0 && len(order[0]) > 0 {
			orderby = " ORDER BY " + order[0]
		}
		sql = sqlParsed + orderby + " OFFSET " + fmt.Sprintf("%v", offset) + " LIMIT " + fmt.Sprintf("%v", limit)
	} else {
		sql = sqlParsed + orderby + constant.MAX_RECORDS
	}

	rows, err := o.QueryRows(sql)

	if err != nil {
		return nil, err
	}

	if limit > 0 {
		pages = math.Ceil(total / limit)
		if page > 1 {
			data["prev"] = page - 1
		} else {
			data["prev"] = page
		}
		if float64(page+1) <= pages {
			data["next"] = page + 1
		} else {
			data["next"] = page
		}
	} else {
		pages = 0
		data["prev"] = 0
		data["next"] = 0
	}

	data["page"] = page
	data["pages"] = pages
	data["total"] = total
	data["rows"] = grid.RowsToInterface(rows)
	data["lines"] = len(data["rows"].([]interface{}))

	return data, nil
}

func (o *Repository) QueryPaginatorWhenGroupBy(columns, table, where, sqlTemplate string, page float64, limit float64, group, order string) (map[string]interface{}, error) {

	var sql string
	var orderby, groupby string
	var total, offset, pages float64
	data := make(map[string]interface{})

	var sqlParsed string
	if len(where) > 0 {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table, where)
	} else {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table)
	}

	if limit == 0 {
		return nil, fmt.Errorf("40020")
	}

	offset = (page - 1) * limit
	var sqlCount string
	if len(group) > 0 {
		sqlCount = "SELECT COUNT(*) FROM (SELECT " + columns + " FROM " + table + " WHERE " + where + " GROUP BY " + group + ")"
	} else {
		sqlCount = "SELECT COUNT(*) FROM " + table + " WHERE " + where
	}
	rowsCount := o.QueryRow(sqlCount)
	rowsCount.Scan(&total)

	if len(group) > 0 {
		groupby = " GROUP BY " + group
	}

	if len(order) > 0 {
		orderby = " ORDER BY " + order
	}

	sql = sqlParsed + groupby + orderby + " OFFSET " + fmt.Sprintf("%v", offset) + " LIMIT " + fmt.Sprintf("%v", limit)

	rows, err := o.QueryRows(sql)

	if err != nil {
		return nil, err
	}

	if limit > 0 {
		pages = math.Ceil(total / limit)
		if page > 1 {
			data["prev"] = page - 1
		} else {
			data["prev"] = page
		}
		if float64(page+1) <= pages {
			data["next"] = page + 1
		} else {
			data["next"] = page
		}
	} else {
		pages = 0
		data["prev"] = 0
		data["next"] = 0
	}

	data["page"] = page
	data["pages"] = pages
	data["total"] = total
	data["rows"] = grid.RowsToInterface(rows)
	data["lines"] = len(data["rows"].([]interface{}))

	return data, nil
}

func (o *Repository) QueryPaginatorWhenGroupByAndHaving(columns, table, where, sqlTemplate string, page float64, limit float64, group, _having_, order string) (map[string]interface{}, error) {

	var sql string
	var orderby, groupby, having string
	var total, offset, pages float64
	data := make(map[string]interface{})

	var sqlParsed string
	if len(where) > 0 {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table, where)
	} else {
		sqlParsed = fmt.Sprintf(sqlTemplate, columns, table)
	}

	if limit == 0 {
		return nil, fmt.Errorf("40020")
	}

	offset = (page - 1) * limit
	var sqlCount string
	if len(group) > 0 {
		if len(_having_) > 0 {
			sqlCount = "SELECT COUNT(*) FROM (SELECT " + columns + " FROM " + table + " WHERE " + where + " GROUP BY " + group + " HAVING" + _having_ + ")"
		} else {
			sqlCount = "SELECT COUNT(*) FROM (SELECT " + columns + " FROM " + table + " WHERE " + where + " GROUP BY " + group + ")"
		}

	} else {
		sqlCount = "SELECT COUNT(*) FROM " + table + " WHERE " + where
	}
	rowsCount := o.QueryRow(sqlCount)
	rowsCount.Scan(&total)

	if len(group) > 0 {
		groupby = " GROUP BY " + group
	}
	if len(_having_) > 0 {
		having = " HAVING" + _having_
	}
	if len(order) > 0 {
		orderby = " ORDER BY " + order
	}

	sql = sqlParsed + groupby + having + orderby + " OFFSET " + fmt.Sprintf("%v", offset) + " LIMIT " + fmt.Sprintf("%v", limit)

	rows, err := o.QueryRows(sql)

	if err != nil {
		return nil, err
	}

	if limit > 0 {
		pages = math.Ceil(total / limit)
		if page > 1 {
			data["prev"] = page - 1
		} else {
			data["prev"] = page
		}
		if float64(page+1) <= pages {
			data["next"] = page + 1
		} else {
			data["next"] = page
		}
	} else {
		pages = 0
		data["prev"] = 0
		data["next"] = 0
	}

	data["page"] = page
	data["pages"] = pages
	data["total"] = total
	data["rows"] = grid.RowsToInterface(rows)
	data["lines"] = len(data["rows"].([]interface{}))

	return data, nil
}

func (o *Repository) WhereCondition(condition interface{}) string {

	var expr string = ""
	cond := condition.([]interface{})

	if len(cond) > 0 {
		expr = cond[0].(string)
		values := cond[1:]
		for _, val := range values {
			expr = strings.Replace(expr, "?", "%v", 1)
			switch reflect.TypeOf(val).Kind().String() {
			case "string":
				expr = fmt.Sprintf(expr, val.(string))
			case "int":
				expr = fmt.Sprintf(expr, val.(int))
			case "int32":
				expr = fmt.Sprintf(expr, val.(int))
			case "int64":
				expr = fmt.Sprintf(expr, val.(int64))
			case "float32":
				expr = fmt.Sprintf(expr, val.(float32))
			case "float64":
				expr = fmt.Sprintf(expr, val.(float64))
			}
		}
	}

	return expr
}
