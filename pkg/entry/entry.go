package entry

type ReportedValue struct {
	Fmt string `json:"fmt"`
}

type Entry struct {
	AsOfDate      string        `json:"asOfDate"`
	PeriodType    string        `json:"periodType"`
	ReportedValue ReportedValue `json:"reportedValue"`
}
