package cluster

type MemberInfo struct {
	ID   string `json:"id"`
	Addr string `json:"address"`
	// SerfPort int    `json:"serf_port"`
	// RaftPort int    `json:"raft_port"`
	// RpcPort  int    `json:"rpc_port"`
	Status string `json:"status"`
}

// func (m *Member) ClusterAddr() *net.TCPAddr {
// 	return &net.TCPAddr{IP: net.ParseIP(m.Addr), Port: m.SerfPort}
// }
//
// func (m *Member) RpcAddr() *net.TCPAddr {
// 	return &net.TCPAddr{IP: net.ParseIP(m.Addr), Port: m.RpcPort}
// }
