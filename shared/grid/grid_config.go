package grid

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"mc-stove/shared/types"
	"sort"
)

type GridParam struct {
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type GridOrder map[string]string

type ExportData struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type GridConfig struct {
	Exec     GridExecType             `json:"exec"`
	Page     string                   `json:"page"`
	RowsPage string                   `json:"rows_page"`
	Params   []map[string][]GridParam `json:"params"`
	Orderby  []GridOrder              `json:"orderby"`
	Options  *GridOptions             `json:"options,omitempty"`
	Export   *ExportData              `json:"export"`
}

type GridExecType uint8

const (
	Filter GridExecType = iota
	Search
)

type GridOptions struct {
	UseSqlQueryPaginator bool
	UseSqlFieldsSearch   bool
}

func NewGridOptions() *GridOptions {
	return &GridOptions{
		UseSqlFieldsSearch:   false,
		UseSqlQueryPaginator: false,
	}
}

func NewGridConfig() *GridConfig {
	return &GridConfig{
		Exec:    Filter,
		Params:  []map[string][]GridParam{},
		Orderby: []GridOrder{},
		Options: &GridOptions{
			UseSqlFieldsSearch:   true,
			UseSqlQueryPaginator: true,
		},
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

func (o *GridConfig) SearchPair(params *Params) []*types.Pair {
	pairs := make([]*types.Pair, 0)
	for _, fieldName := range params.listSearchFields {
		for _, searchParam := range params.listSearchValue {
			pairs = append(pairs, &types.Pair{Key: fieldName, Value: searchParam.Value})
		}
	}
	return pairs
}

func (o *GridConfig) Hash(params *Params) (string, error) {
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

	b, err := json.Marshal(compose)
	if err != nil {
		return "", err
	}
	h := sha256.New()
	h.Write(b)

	return hex.EncodeToString(h.Sum(nil)), err
}

func (o *GridConfig) UseSqlQueryPaginator() bool {
	return o.Options != nil && o.Options.UseSqlQueryPaginator
}

func (o *GridConfig) UseSqlFieldsSearch() bool {
	return o.Options != nil && o.Options.UseSqlFieldsSearch
}
