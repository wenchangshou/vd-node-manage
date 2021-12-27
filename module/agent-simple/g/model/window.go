package model

type WindowStatus int

const (
	INIT WindowStatus = iota
	OPENING
	RUNNING
	CLOSE
	ABNORMAL
)

func (w WindowStatus) String() string {
	return [...]string{"Init", "Running", "Opening", "Close", "Abnormal"}[w]
}

type WindowStyle struct {
	WindowStyle string `json:"WindowStyle"`
}

type Window struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
	Z      int `json:"z"`
}
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
