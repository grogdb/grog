package cluster

import (
	"net"
)

type Config struct {
	DataDir        string
	NodeID         string
	ClusterAddr    *net.TCPAddr
	ClusterRPCAddr *net.TCPAddr
	GossipAddr     *net.TCPAddr
	HTTPAddr       *net.TCPAddr
}
