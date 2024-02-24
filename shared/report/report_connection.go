package report

import "gorm.io/gorm"

type ReportConnection struct {
	HotDb  *gorm.DB
	ColdDb *gorm.DB
}

func NewReportConnection() *ReportConnection {
	return &ReportConnection{}
}
