package e

type ResourceSource struct {
	Type  string `json:"Type"`
	Fname string `json:"Fname"`
	URL   string `json:"URL"`
	RID   string `json:"RID"`
}
type ResourceArgument struct {
}
type WindowStyle struct {
	WindowStyle string `json:"WindowStyle"`
}

//func (style WindowStyle) toString(){
//	if
//tart
type LayoutArgument struct {
	LayoutID string      `json:"LayoutId"`
	Style    WindowStyle `json:"Style"`
	Windows  []Window    `json:"Windows"`
	Kill     bool        `json:"Kill"`
}

type OpenLayoutForm struct {
	Action    string         `json:"Action"`
	Arguments LayoutArgument `json:"Arguments"`
}
