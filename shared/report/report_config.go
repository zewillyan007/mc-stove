package report

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"mc-stove/shared/types"
	"sort"
)

type ReportParam struct {
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type ReportOrder map[string]string

type ReportGroup []string

type ReportConfig struct {
	Name        string                     `json:"-"`
	Page        string                     `json:"page"`
	RowsPage    string                     `json:"rows_page"`
	Params      []map[string][]ReportParam `json:"params"`
	Orderby     []ReportOrder              `json:"orderby"`
	Groupby     ReportGroup                `json:"groupby"`
	Inputs      map[string]interface{}     `json:"inputs"`
	FilteredBy  map[string]interface{}     `json:"filtered_by"`
	Options     *ReportOptions             `json:"options,omitempty"`
	RecordLimit int                        `json:"record_limit"`
}

type ReportOptions struct {
	UseSqlQueryPaginator bool
	UseSqlFieldsSearch   bool
}

func NewReportOptions() *ReportOptions {
	return &ReportOptions{
		UseSqlFieldsSearch:   false,
		UseSqlQueryPaginator: false,
	}
}

func NewReportConfig() *ReportConfig {

	return &ReportConfig{
		Name:       "",
		Params:     []map[string][]ReportParam{},
		Orderby:    []ReportOrder{},
		Groupby:    ReportGroup{},
		Inputs:     map[string]interface{}{},
		FilteredBy: map[string]interface{}{},
		Options: &ReportOptions{
			UseSqlFieldsSearch:   true,
			UseSqlQueryPaginator: true,
		},
		RecordLimit: 0,
	}
}

type hashRow struct {
	FieldName string
	Options   []*hashFieldOptions
}

type hashFieldOptions struct {
	Operator string
	Value    interface{}
}

type hashCompose struct {
	Fields []*hashRow
}

func (o *ReportConfig) SearchPair(params *Params) []*types.Pair {
	pairs := make([]*types.Pair, 0)
	for _, fieldName := range params.listSearchFields {
		for _, searchParam := range params.listSearchValue {
			pairs = append(pairs, &types.Pair{Key: fieldName, Value: searchParam.Value})
		}
	}
	return pairs
}

func (o *ReportConfig) Hash(params *Params) (string, error) {
	names := params.GetListKeys()
	sort.Strings(names)
	compose := &hashCompose{Fields: make([]*hashRow, len(names))}
	for n, name := range names {
		compose.Fields[n] = &hashRow{FieldName: name}
		options := params.GetParamValue(name)
		compose.Fields[n].Options = make([]*hashFieldOptions, len(options))
		for ix, m := range options {
			compose.Fields[n].Options[ix] = &hashFieldOptions{
				Operator: m["operator"].(string),
				Value:    m["value"],
			}
		}
	}

	if o.Orderby != nil && len(o.Orderby) > 0 {
		names = make([]string, len(o.Orderby))
		orders := make(map[string]string)

		for n, order := range o.Orderby {
			for k, v := range order {
				names[n] = k
				orders[k] = v
			}
		}

		sort.Strings(names)
		row := &hashRow{FieldName: "orders"}
		compose.Fields = append(compose.Fields, row)
		row.Options = make([]*hashFieldOptions, len(names))

		for n, name := range names {
			row.Options[n] = &hashFieldOptions{
				Operator: name,
				Value:    orders[name],
			}
		}
	}

	if o.Groupby != nil && len(o.Groupby) > 0 {
		sort.Strings(o.Groupby)
		row := &hashRow{FieldName: "groupby"}
		compose.Fields = append(compose.Fields, row)
		row.Options = make([]*hashFieldOptions, len(o.Groupby))

		for n, name := range o.Groupby {
			row.Options[n] = &hashFieldOptions{
				Operator: "group",
				Value:    name,
			}
		}
	}

	if alert, ok := params.GetList()["alert"]; ok {
		names = make([]string, len(alert))

		for n, item := range alert {
			names[n] = item["value"].(string)
		}

		sort.Strings(names)

		row := &hashRow{FieldName: "alerts"}
		compose.Fields = append(compose.Fields, row)
		row.Options = make([]*hashFieldOptions, len(alert))

		for n, name := range names {
			row.Options[n] = &hashFieldOptions{
				Operator: "=",
				Value:    name,
			}
		}
	}

	b, err := json.Marshal(compose)
	if err != nil {
		return "", err
	}
	h := sha256.New()
	h.Write(b)

	return hex.EncodeToString(h.Sum(nil)), err
}

func (o *ReportConfig) UseSqlQueryPaginator() bool {
	return o.Options != nil && o.Options.UseSqlQueryPaginator
}

func (o *ReportConfig) UseSqlFieldsSearch() bool {
	return o.Options != nil && o.Options.UseSqlFieldsSearch
}
