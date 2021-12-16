package process

type IProcess interface {
	StartProcessAsCurrentUser(appPath, cmdLine, workDir string, backstage bool) (uint32, error)
}
