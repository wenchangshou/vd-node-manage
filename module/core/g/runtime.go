package g

// RunTimeParams 运行时参数
type RunTimeParams struct {
	Register bool   `json:"register"`
	ID       string `json:"id"`
}

var (
	GRunTimeParams *RunTimeParams = &RunTimeParams{
		Register: false,
		ID:       "",
	}
)
