package libcontainer

import criu "github.com/checkpoint-restore/go-criu/v4/rpc"

type CriuLogLevel int32

const (
	CRIU_LOG_DEFAULT CriuLogLevel = 0 + iota
	CRIU_LOG_MSG
	CRIU_LOG_ERROR
	CRIU_LOG_WARN
	CRIU_LOG_INFO
	CRIU_LOG_DEBUG
)

type CriuPageServerInfo struct {
	Address string // IP address of CRIU page server
	Port    int32  // port number of CRIU page server
}

type VethPairName struct {
	ContainerInterfaceName string
	HostInterfaceName      string
}

type CriuOpts struct {
	ImagesDirectory         string             // directory for storing image files
	WorkDirectory           string             // directory to cd and write logs/pidfiles/stats to
	ParentImage             string             // directory for storing parent image files in pre-dump and dump
	LeaveRunning            bool               // leave container in running state after checkpoint
	TcpEstablished          bool               // checkpoint/restore established TCP connections
	ExternalUnixConnections bool               // allow external unix connections
	ShellJob                bool               // allow to dump and restore shell jobs
	FileLocks               bool               // handle file locks, for safety
	PreDump                 bool               // call criu predump to perform iterative checkpoint
	PageServer              CriuPageServerInfo // allow to dump to criu page server
	VethPairs               []VethPairName     // pass the veth to criu when restore
	ManageCgroupsMode       criu.CriuCgMode    // dump or restore cgroup mode
	EmptyNs                 uint32             // don't c/r properties for namespace from this mask
	AutoDedup               bool               // auto deduplication for incremental dumps
	LazyPages               bool               // restore memory pages lazily using userfaultfd
	StatusFd                int                // fd for feedback when lazy server is ready
	External                []string           // list of external devices
	LogLevel                CriuLogLevel       // log verbosity level
}
