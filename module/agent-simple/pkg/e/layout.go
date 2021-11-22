package e

type Win struct {
	WinID     string
	X         int
	Y         int
	Width     int
	Height    int
	Z         int
	Type      string
	Service   string
	FName     string `json:"fName"`
	RID       string `json:"RID"`
	ProcessID int
	IsOpen    bool
	Status    WindowStatus
	Style     WindowStyle
	Arguments interface{}
	URL       string
}
