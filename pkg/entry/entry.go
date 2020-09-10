package entry

import "github.com/jinzhu/gorm"

type ReportedValue struct {
	Fmt string `json:"fmt"`
}

type Entry struct {
	gorm.Model

	AsOfDate      string        `json:"asOfDate"`
	PeriodType    string        `json:"periodType"`
	ReportedValue ReportedValue `json:"reportedValue"`
}
