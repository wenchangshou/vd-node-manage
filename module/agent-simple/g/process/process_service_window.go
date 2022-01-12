package process

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type WindowServiceProcess struct {
}
type PsInfo struct {
	ProcessID       int    //进程id
	ParentProcessID int    //上级进程id
	ProcessName     string //进程名称
}
type PsTree map[int]PsInfo

func (t PsTree) findDescendants(pid int) []int {
	var children []int
	for _, ps := range t {
		if ps.ParentProcessID == pid {
			children = append(children, ps.ProcessID)
			children = append(children, t.findDescendants(ps.ProcessID)...)
		}
	}
	return children
}

// SnapshotSysProcesses 获取当前所有的进程的信息
func SnapshotSysProcesses() (PsTree, error) {
	ss, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(ss)
	ps := make(PsTree)
	var pe syscall.ProcessEntry32
	pe.Size = uint32(unsafe.Sizeof(pe))
	if err = syscall.Process32First(ss, &pe); err != nil {
		return nil, err
	}
	for {
		tmp := PsInfo{
			ProcessID:       int(pe.ProcessID),
			ParentProcessID: int(pe.ParentProcessID),
			ProcessName:     windows.UTF16ToString(pe.ExeFile[:]),
		}
		ps[int(pe.ProcessID)] = tmp
		err = syscall.Process32Next(ss, &pe)
		if err == syscall.ERROR_NO_MORE_FILES {
			return ps, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
func CheckThreadExists(id uint32) bool {
	tree, _ := SnapshotSysProcesses()
	for _, ps := range tree {
		if ps.ProcessID == int(id) {
			return true
		}
	}
	return false
}
func (w WindowServiceProcess) GetThreadStatus(id uint32) bool {
	return CheckThreadExists(id)
}

var (
	modwtsapi  *windows.LazyDLL = windows.NewLazySystemDLL("wtsapi32.dll")
	modkernel  *windows.LazyDLL = windows.NewLazySystemDLL("kernel32.dll")
	modadvapi  *windows.LazyDLL = windows.NewLazySystemDLL("advapi32.dll")
	moduserenv *windows.LazyDLL = windows.NewLazySystemDLL("userenv.dll")

	procWTSEnumerateSessionsW        *windows.LazyProc = modwtsapi.NewProc("WTSEnumerateSessionsW")
	procWTSGetActiveConsoleSessionId *windows.LazyProc = modkernel.NewProc("WTSGetActiveConsoleSessionId")
	procWTSQueryUserToken            *windows.LazyProc = modwtsapi.NewProc("WTSQueryUserToken")
	procDuplicateTokenEx             *windows.LazyProc = modadvapi.NewProc("DuplicateTokenEx")
	procCreateEnvironmentBlock       *windows.LazyProc = moduserenv.NewProc("CreateEnvironmentBlock")
	procCreateProcessAsUser          *windows.LazyProc = modadvapi.NewProc("CreateProcessAsUserW")
)

const (
	WTS_CURRENT_SERVER_HANDLE uintptr = 0
)

type WTS_CONNECTSTATE_CLASS int

const (
	WTSActive WTS_CONNECTSTATE_CLASS = iota
	WTSCOnnected
	WTSConnectQuery
	WTSShadow
	WTSDisconnected
	WTSIdle
	WTSListen
	WTSReset
	WTSDown
	WTSInit
)
const (
	PROCESS_QUERY_INFORMATION = 1 << 10
	PROCESS_VM_READ           = 1 << 4
)

type PROCESS_MEMORY_COUNTERS struct {
	cb                         uint32
	PageFaultCount             uint32
	PeakWorkingSetSize         uint64
	WorkingSetSize             uint64
	QuotaPeakPagedPoolUsage    uint64
	QuotaPagedPoolUsage        uint64
	QuotaPeakNonPagedPoolUsage uint64
	QuotaNonPagedPoolUsage     uint64
	PagefileUsage              uint64
	PeakPagefileUsage          uint64
}

type SECURITY_IMPERSONATION_LEVEL int

const (
	SecurityAnonymous SECURITY_IMPERSONATION_LEVEL = iota
	SecurityIdentification
	SecurityImpersonation
	SecurityDelegation
)

type TOKEN_TYPE int

const (
	TokenPrimary TOKEN_TYPE = iota + 1
	TokenImpersonazion
)

type SW int

const (
	SW_HIDE            SW = 0
	SW_SHOWNORMAL         = 1
	SW_NORMAL             = 1
	SW_SHOWMINIMIZED      = 2
	SW_SHOWMAXIMIZED      = 3
	SW_MAXIMIZE           = 3
	SW_SHOWNOACTIVATE     = 4
	SW_SHOW               = 5
	SW_MINIMIZE           = 6
	SW_SHOWMINNOACTIVE    = 7
	SW_SHOWNA             = 8
	SW_RESTORE            = 9
	SW_SHOWDEFAULT        = 10
	SW_MAX                = 1
)

type WTS_SESSION_INFO struct {
	SessionID      windows.Handle
	WinStationName *uint16
	State          WTS_CONNECTSTATE_CLASS
}

const (
	CREATE_UNICODE_ENVIRONMENT uint16 = 0x00000400
	CREATE_NO_WINDOW                  = 0x08000000
	CREATE_NEW_CONSOLE                = 0x00000010
)

func WTSEnumerateSessions() ([]*WTS_SESSION_INFO, error) {
	var (
		sessionInformation windows.Handle      = windows.Handle(0)
		sessionCount       int                 = 0
		SessionList        []*WTS_SESSION_INFO = make([]*WTS_SESSION_INFO, 0)
	)
	if returnCode, _, err := procWTSEnumerateSessionsW.Call(WTS_CURRENT_SERVER_HANDLE, 0, 1, uintptr(unsafe.Pointer(&sessionInformation)), uintptr(unsafe.Pointer(&sessionCount))); returnCode == 0 {
		return nil, fmt.Errorf("call native WTSEnumerateSessionsW: %s", err)
	}
	structSize := unsafe.Sizeof(WTS_SESSION_INFO{})
	current := uintptr(sessionInformation)
	for i := 0; i < sessionCount; i++ {
		SessionList = append(SessionList, (*WTS_SESSION_INFO)(unsafe.Pointer(current)))
		current += structSize
	}
	return SessionList, nil
}
func GetCurrentUserSessionId() (windows.Handle, error) {
	sessionList, err := WTSEnumerateSessions()
	if err != nil {
		return 0xFFFFFFFF, fmt.Errorf("get current user session token:%s", err)
	}
	for i := range sessionList {
		if sessionList[i].State == WTSActive {
			return sessionList[i].SessionID, nil
		}
	}
	if sessionId, _, err := procWTSGetActiveConsoleSessionId.Call(); sessionId == 0xFFFFFFFF {
		return 0xFFFFFFFF, fmt.Errorf("get current user session token:call native WTSGetActiveConsoleSessionId:%s", err)
	} else {
		return windows.Handle(sessionId), nil
	}
}
func (w WindowServiceProcess) StartProcessAsCurrentUser(appPath, cmdLine, workDir string, backstage bool) (uint32, error) {
	var (
		sessionId windows.Handle
		userToken windows.Token
		envInfo   windows.Handle

		startupInfo windows.StartupInfo
		processInfo windows.ProcessInformation
		commandLine uintptr = 0
		workingDir  uintptr = 0
		err         error
	)
	if sessionId, err = GetCurrentUserSessionId(); err != nil {
		return 0, err
	}
	if userToken, err = DuplicateUserTokenFromSessionID(sessionId); err != nil {
		return 0, fmt.Errorf("get duplicate user token for current user session:%s", err)
	}
	if returnCode, _, err := procCreateEnvironmentBlock.Call(uintptr(unsafe.Pointer(&envInfo)), uintptr(userToken), 0); returnCode == 0 {
		return 0, fmt.Errorf("create environment details for process:%s", err)
	}
	if backstage {
		startupInfo.ShowWindow = uint16(SW_HIDE)
		startupInfo.Flags = windows.STARTF_USESHOWWINDOW | windows.STARTF_USESHOWWINDOW
	} else {
		startupInfo.ShowWindow = SW_SHOW
	}
	creationFlags := CREATE_UNICODE_ENVIRONMENT | CREATE_NEW_CONSOLE
	//creationFlags:=CREATE_NO_WINDOW
	startupInfo.Desktop = windows.StringToUTF16Ptr("winsta0\\default")
	if len(cmdLine) > 0 {
		commandLine = uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(cmdLine)))
	}
	if len(workDir) > 0 {
		workingDir = uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(workDir)))
	}
	if returnCode, _, err := procCreateProcessAsUser.Call(
		uintptr(userToken), uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(appPath))), commandLine, 0, 0, 0,
		uintptr(creationFlags), uintptr(envInfo), workingDir, uintptr(unsafe.Pointer(&startupInfo)), uintptr(unsafe.Pointer(&processInfo)),
	); returnCode == 0 {
		return 0, fmt.Errorf("create process as user:%s", err)
	}
	return processInfo.ProcessId, nil
}

// DuplicateUserTokenFromSessionID will attempt
// to duplicate the user token for the user logged
// into the provided session ID
func DuplicateUserTokenFromSessionID(sessionId windows.Handle) (windows.Token, error) {
	var (
		impersonationToken windows.Handle = 0
		userToken          windows.Token  = 0
	)

	if returnCode, _, err := procWTSQueryUserToken.Call(uintptr(sessionId), uintptr(unsafe.Pointer(&impersonationToken))); returnCode == 0 {
		return 0xFFFFFFFF, fmt.Errorf("call native WTSQueryUserToken: %s", err)
	}

	if returnCode, _, err := procDuplicateTokenEx.Call(uintptr(impersonationToken), 0, 0, uintptr(SecurityImpersonation), uintptr(TokenPrimary), uintptr(unsafe.Pointer(&userToken))); returnCode == 0 {
		return 0xFFFFFFFF, fmt.Errorf("call native DuplicateTokenEx: %s", err)
	}

	if err := windows.CloseHandle(impersonationToken); err != nil {
		return 0xFFFFFFFF, fmt.Errorf("close windows handle used for token duplication: %s", err)
	}

	return userToken, nil
}
func NewWindowServiceProcess() IProcess {
	return &WindowServiceProcess{}
}
