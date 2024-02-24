package report

import (
	"strings"
)

type Groups struct {
	list []string
}

func NewGroups() *Groups {
	return &Groups{
		list: []string{},
	}
}

func (o *Groups) Validate(columns []string) []string {

	var invalidFields []string
	filterKeys := o.GetList()
	columnsString := strings.Join(columns, ",")

	for _, filter := range filterKeys {
		if !strings.Contains(columnsString, filter) {
			invalidFields = append(invalidFields, filter)
		}
	}
	return invalidFields
}

func (o *Groups) GetList() []string {
	return o.list
}

func (o *Groups) LoadReportGroups(report *ReportConfig) {
	for _, v := range report.Groupby {
		o.Add(strings.ToLower(v))
	}
}

func (o *Groups) ToString() string {
	return strings.Join(o.list, ", ")
}

func (o *Groups) Add(field string) {
	o.list = append(o.list, field)
}
