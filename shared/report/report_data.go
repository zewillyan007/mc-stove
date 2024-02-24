package report

type ReportData struct {
	Lines int
	Next  int
	Page  int
	Pages int
	Prev  int
	Rows  []interface{}
	Total int
}

func NewReportData() *ReportData {
	return &ReportData{
		Lines: 0,
		Next:  1,
		Page:  1,
		Pages: 1,
		Prev:  1,
		Rows:  []interface{}{},
		Total: 0,
	}
}
