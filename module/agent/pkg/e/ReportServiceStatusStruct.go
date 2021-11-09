package e
type WindowLayoutInfo struct{
	Open bool
	Wid string
}
type ReportServiceStatusStruct struct{
	LayoutId string
	Windows []WindowLayoutInfo
}
