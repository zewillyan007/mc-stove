package usecase

import (
	"mc-stove/plantation/core/port"
	"mc-stove/shared/grid"
	"strconv"
	"strings"
)

type PlantUseCaseGrid struct {
	grid.Grid
	repository port.PlantIRepository
}

func NewPlantUseCaseGrid(repository port.PlantIRepository) *PlantUseCaseGrid {
	return &PlantUseCaseGrid{repository: repository}
}

func (o *PlantUseCaseGrid) table() string {

	return "plantation.plant"
}

func (o *PlantUseCaseGrid) columns() []string {

	return []string{
		"id",
		"number",
		"length",
		"width",
		"height",
	}
}

func (o *PlantUseCaseGrid) mandatory() []string {

	return []string{}
}

func (o *PlantUseCaseGrid) searchFields() map[string]string {

	return map[string]string{
		"number": "numeric",
		"length": "numeric",
		"width":  "numeric",
		"height": "numeric",
	}
}

func (o *PlantUseCaseGrid) orderFields() map[string]string {

	return map[string]string{
		"number": "numeric",
		"length": "numeric",
		"width":  "numeric",
		"height": "numeric",
	}
}

func (o *PlantUseCaseGrid) Execute(GridConfig *grid.GridConfig) (data map[string]interface{}, err error) {

	var sql string = ""
	var order string = ""
	var page, limit float64
	var where []string = []string{}

	prepare, err := o.Prepare(GridConfig, o.columns(), o.mandatory(), o.searchFields())
	if err != nil {
		return nil, err
	}
	params := prepare["params"].(*grid.Params)
	orders := prepare["orders"].(*grid.Orders)

	// where = append(where, "(status <> 'DELETED')")

	if len(params.ToString()) > 0 {
		where = append(where, params.ToString())
	}

	if GridConfig.UseSqlFieldsSearch() && len(params.ToStringSearch()) > 0 {
		where = append(where, "("+params.ToStringSearch()+")")
	}

	if len(orders.GetList()) > 0 {
		order = orders.ToStringTranslate(o.orderFields())
	}

	sql = "SELECT %s FROM %s"
	if len(where) > 0 {
		sql = sql + " WHERE %s"
	}

	page, _ = strconv.ParseFloat(GridConfig.Page, 64)
	limit, _ = strconv.ParseFloat(GridConfig.RowsPage, 64)

	if GridConfig.UseSqlQueryPaginator() {
		data, err = o.repository.SqlQueryPaginator(strings.Join(o.columns(), ","), o.table(), strings.Join(where, " AND "), sql, page, limit, order)
	} else {
		data, err = o.repository.SqlQueryData(strings.Join(o.columns(), ","), o.table(), strings.Join(where, " AND "), sql, order)
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}
