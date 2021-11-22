package IService

type ComputerService interface {
	ReportServiceInfo(id string, ip string, mac string, name string) error
	Report() error
	IsRegister() (bool, error)
	AddComputerProject(projectReleaseID string) error
	DeleteComputerProject(projectReleaseID string) error
	DeleteComputerResource(resourceID string) error
	AddComputerResource(resourceID string) error
	Heartbeat() error
}
