package IService

type DeviceService interface {
	ReportServiceInfo(id uint, ip string, mac string, name string) error
	Report() error
	IsRegister() (bool, error)
	DeleteComputerResource(resourceID uint) error
	AddComputerResource(resourceID uint) error
	Heartbeat() error
	AddComputerProject(id uint) error
	DeleteComputerProject(id uint) error
}
