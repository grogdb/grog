package bootstrap

import (
	"fmt"
	"github.com/rs/zerolog"
)

var rootLogger *zerolog.Logger

const (
	flagAdvertiseAddr     = "advertise-addr"
	flagAdvertiseAddrDesc = "Address advertised to other nodes in the cluster (default: first available private IP)"
	flagBindAddr          = "bind-addr"
	flagBindAddrDesc      = "Address to bind services to"
	flagClusterPort       = "cluster-port"
	flagClusterPortDesc   = "Port used for nodes communication"
	flagDataDir           = "data-dir"
	flagDataDirDesc       = "Data directory"
	flagDebug             = "debug"
	flagDebugShort        = "D"
	flagDebugDesc         = "Enable debug logging"
	flagEncrypt           = "encrypt"
	flagEncryptDesc       = "Set the gossip encryption key"
	flagGossipPort        = "gossip-port"
	flagGossipPortDesc    = "Port used for cluster management"
	flagJoin              = "join"
	flagJoinDesc          = "Address of a node to join a cluster. Can be specified as a comma separated list"
	flagLogFormat         = "log-format"
	flagLogFormatDesc     = "Set the log format [console, json]"
	flagServerAddr        = "addr"
	flagServerAddrDesc    = "Server address"
	flagVerbose           = "verbose"
	flagVerboseDesc       = "Enable verbose logging"
)

const (
	defaultBindAddr    = "0.0.0.0"
	defaultClusterPort = 44420 // cluster RPC
	defaultGossipPort  = 44421 // serf
	defaultDataDir     = "~/.grog"
	defaultRPCPort     = 6666
	defaultHTTPPort    = 8888
	logFormatConsole   = "console"
	logFormatJSON      = "json"
)

var defaultServerAddr = fmt.Sprintf("127.0.0.1:%d", defaultRPCPort)
