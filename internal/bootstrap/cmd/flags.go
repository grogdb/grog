package cmd

import (
	"fmt"

	"github.com/rs/zerolog"
)

var rootLogger *zerolog.Logger

const (
	flagAdvertiseAddr     = "advertise-addr"
	flagAdvertiseAddrDesc = "address advertised to other nodes in the cluster (default: first available private IP)"
	flagBindAddr          = "bind-addr"
	flagBindAddrDesc      = "address to bind services to"
	flagClusterPort       = "cluster-port"
	flagClusterPortDesc   = "port used for nodes communication"
	flagDataDir           = "data-dir"
	flagDataDirDesc       = "data directory"
	flagDebug             = "debug"
	flagDebugShort        = "d"
	flagDebugDesc         = "enable debug logging"
	flagEncrypt           = "encrypt"
	flagEncryptDesc       = "set the gossip encryption key"
	flagFile              = "file"
	flagFileShort         = "f"
	flagFileDesc          = "source file"
	flagGossipPort        = "gossip-port"
	flagGossipPortDesc    = "port used for cluster management"
	flagJoin              = "join"
	flagJoinDesc          = "address of a node to join a cluster. Can be specified as a comma separated list"
	flagLogFormat         = "log-format"
	flagLogFormatDesc     = "set the log format [console, json]"
	flagServerAddr        = "addr"
	flagServerAddrDesc    = "server address"
	flagVerbose           = "verbose"
	flagVerboseDesc       = "enable verbose logging"
)

const (
	defaultBindAddr    = "0.0.0.0"
	defaultClusterPort = 4420 // cluster RPC
	defaultGossipPort  = 4421 // serf
	defaultDataDir     = "~/.grog"
	defaultRPCPort     = 6666
	defaultHTTPPort    = 8888
	logFormatConsole   = "console"
	logFormatJSON      = "json"
)

var defaultServerAddr = fmt.Sprintf("127.0.0.1:%d", defaultRPCPort)
