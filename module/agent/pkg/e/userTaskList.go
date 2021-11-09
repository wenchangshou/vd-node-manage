package e

type ProjestRelease struct {
	ID        uint   `json:"ID"`
	Number    string `json:"number"`
	Content   string `json:"content"`
	Mode      string `json:"mode"`
	Depend    int    `json:"depend"`
	Arguments string `json:"arguments"`
	Control   string `json:"control"`
	File      File
	Project   Project
}

type Project struct {
	ID          int
	Start       string
	Name        string
	Category    string
	Description string
	Arguments   string
	Control     string
}
type File struct {
	ID         int
	Name       string `json:"name"`
	Mode       string `json:"mode"`
	SourceName string
	UserId     int
	Size       int
}
type Task struct {
	ID       int                    `json:"ID" mapstructure:"ID"`
	Options  map[string]interface{} `mapstructure:",remain"`
	Action   int
	Depend   int
	Schedule int
	Active   bool
}
