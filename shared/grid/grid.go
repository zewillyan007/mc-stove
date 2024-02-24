package grid

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"mc-stove/shared/types"
	"mc-stove/shared/util"
	"strings"
)

type Grid struct {
}

func (o *Grid) Prepare(GridConfig *GridConfig, columns, mandatory []string, search map[string]string) (map[string]interface{}, error) {

	prep := make(map[string]interface{}, 0)
	params := NewParams()
	orders := NewOrders()
	params.SetSearchFields(search)
	params.LoadGridParams(GridConfig)
	orders.LoadGridOrders(GridConfig)
	invalidFields := make(map[string][]string, 0)

	mandatoryParams := params.ValidateMandatory(mandatory)
	if len(mandatoryParams) > 0 {
		invalidFields["mandatory"] = mandatoryParams
	}

	invalidParams := params.Validate(columns)
	if len(invalidParams) > 0 {
		invalidFields["params"] = invalidParams
	}

	invalidOrders := orders.Validate(columns)
	if len(invalidOrders) > 0 {
		invalidFields["orders"] = invalidOrders
	}

	if len(invalidFields) > 0 {
		msg := make(map[string]string, 0)
		if len(invalidFields["mandatory"]) > 0 {
			msg["mandatory"] = strings.Join(invalidFields["mandatory"], ", ")
		}
		if len(invalidFields["params"]) > 0 {
			msg["params"] = strings.Join(invalidFields["params"], ", ")
		}
		if len(invalidFields["orders"]) > 0 {
			msg["orders"] = strings.Join(invalidFields["orders"], ", ")
		}
		_json_, _ := json.Marshal(msg)
		return nil, errors.New(string(_json_))
	}

	prep["params"] = params
	prep["orders"] = orders

	return prep, nil
}

func (o *Grid) ParamIntervalDate(params *Params, nameInitDate, nameFinDate string) map[string]interface{} {

	mapa := make(map[string]interface{})
	paramsNew := NewParams()
	listParams := params.GetList()
	for nameParam, valueParam := range listParams {
		if nameParam != nameInitDate && nameParam != nameFinDate {
			for _, val := range valueParam {
				paramsNew.Add(nameParam, val["operator"].(string), val["value"])
			}
		} else {
			if nameParam == nameInitDate {
				mapa[nameInitDate] = valueParam[0]["value"].(string)
			}
			if nameParam == nameFinDate {
				mapa[nameFinDate] = valueParam[0]["value"].(string)
			}
		}
	}
	mapa["params"] = paramsNew
	return mapa
}

func (o *Grid) ParamExtract(params *Params, paramName string) map[string]interface{} {

	mapa := make(map[string]interface{})
	listParams := params.GetList()
	params.ClearList()
	for nameParam, valueParam := range listParams {

		if nameParam != paramName {
			for _, val := range valueParam {
				params.Add(nameParam, val["operator"].(string), val["value"])
			}
		} else {
			mapa[paramName] = valueParam
		}
	}
	mapa["params"] = params
	return mapa
}

func (o *Grid) VisionCompaniesAndRegionalsArrayToString(params *Params, companies, regionals []int64) (string, string) {

	var strCompanies string
	if len(companies) > 0 {
		strCompanies = fmt.Sprintf("(id_company in (%s))", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(companies)), ","), "[]"))
	}

	var strRegionals string
	if len(regionals) > 0 {
		strRegionals = fmt.Sprintf("(id_regional in (%s))", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(regionals)), ","), "[]"))
	}

	return strCompanies, strRegionals
}

func (o *Grid) ArrayToString(params *Params, fieldName string, array []int64) string {

	paramsList := params.GetList()

	var out string
	if len(array) > 0 {
		if _, found := paramsList[fieldName]; !found {
			out = fmt.Sprintf("("+fieldName+"in (%s))", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ","), "[]"))
		}
	}

	return out
}

func (o *Grid) Paginate(data map[string]interface{}, page, limit float64) {
	rows := data["rows"].([]interface{})
	total := float64(len(rows))
	var pages float64
	info := &types.CacheSubsetInfo{
		Preview: 0,
		Next:    0,
	}

	if limit > 0 {
		offset := (page - 1) * limit
		pages = math.Ceil(total / limit)
		if page > 1 {
			info.Preview = int(page - 1)
		} else {
			info.Preview = int(page)
		}
		if page+1 <= pages {
			info.Next = int(page + 1)
		} else {
			info.Next = int(page)
		}
		limit = limit + offset
		if limit > total {
			limit = total
		}
		if offset > limit {
			rows = []interface{}{}
		} else {
			rows = rows[int(offset):int(limit)]
		}
	} else {
		pages = 0
	}

	info.Page = int(page)
	info.Pages = int(pages)
	info.Total = int(total)
	info.Lines = len(rows)

	data["rows"] = rows
	o.SetPaginateInfo(data, info)
}

func (o *Grid) SetPaginateInfo(data map[string]interface{}, info *types.CacheSubsetInfo) {
	data["prev"] = info.Preview
	data["next"] = info.Next
	data["page"] = info.Page
	data["pages"] = info.Pages
	data["total"] = info.Total
	data["lines"] = info.Lines
}

func (o *Grid) ApplySearch(data map[string]interface{}, gridConfig *GridConfig) {
	params := NewParams()
	params.LoadGridParams(gridConfig)

	if len(params.listSearchFields) > 0 && len(params.listSearchValue) > 0 {
		searchData := make([]interface{}, 0)

		if rows, ok := data["rows"].([]interface{}); ok {
			for _, row := range rows {
				m := row.(map[string]interface{})

				for _, fieldName := range params.listSearchFields {
					if fieldValue, ok := m[fieldName]; ok {
						strFieldValue := util.InterfaceToString(fieldValue)
						for _, searchParam := range params.listSearchValue {
							if strings.Contains(strings.ToLower(strFieldValue), strings.ToLower(util.InterfaceToString(searchParam.Value))) {
								searchData = append(searchData, row)
								break
							}
						}
					}
				}
			}
		}

		data["rows"] = searchData
	}
}

func (o *Grid) DataToStructDataGrid(data map[string]interface{}, struct_ interface{}) (interface{}, error) {
	jsonStr, err := json.Marshal(data)

	if err == nil {
		err = json.Unmarshal(jsonStr, struct_)
	}

	return struct_, err
}
