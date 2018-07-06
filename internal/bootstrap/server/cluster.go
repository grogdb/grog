package server

func CommandFlags() {

}

func NewManager() (Manager, error) {
	return &clusterManager{}, nil
}

type Manager interface {
	Start(shutdownCh <-chan struct{}, errCh chan error)
	// Encrypted() bool
	// GetLANCoordinate() (lib.CoordinateSet, error)
	// Leave() error
	// LANMembers() []serf.Member
	// LANMembersAllSegments() ([]serf.Member, error)
	// LANSegmentMembers(segment string) ([]serf.Member, error)
	// LocalMember() serf.Member
	// JoinLAN(addrs []string) (n int, err error)
	// RemoveFailedNode(node string) error
	// RPC(method string, args interface{}, reply interface{}) error
	// SnapshotRPC(args *structs.SnapshotRequest, in io.Reader, out io.Writer, replyFn structs.SnapshotReplyFn) error
	// Shutdown() error
	// Stats() map[string]map[string]string
	// ReloadConfig(config *consul.Config) error
	// enterpriseDelegate
}

type clusterManager struct {
}

func (m *clusterManager) Start(shutdownCh <-chan struct{}, errCh chan error) {

}
