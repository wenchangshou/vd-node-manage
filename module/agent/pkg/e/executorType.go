package e
type ExecuteType uint
const (
	InstallProjectAction ExecuteType = iota
	InstallResourceAction
	UpgradeProjectAction
	DeleteResource
	DeleteProject
)