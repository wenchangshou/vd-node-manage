package process

type Process interface {
	Pid() int
	PPid() int
	Executable() string
}
