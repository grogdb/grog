package cmd

import (
	"fmt"

	"github.com/rs/zerolog"
)

var rootLogger *zerolog.Logger

const (
	flagAdvertiseAddr       = "advertise-addr"
	flagAdvertiseAddrDesc   = "address advertised to other nodes in the cluster (default: private IP)"
	flagBindAddr            = "bind-addr"
	flagBindAddrDesc        = "address to bind services to"
	flagBootstrapExpect     = "bootstrap-expect"
	flagBootstrapExpectDesc = "set the minimum amount of nodes required to boostrap the cluster"
	flagClientHTTPPort      = "client-http-port"
	flagClientHTTPPortDesc  = "port used for client HTTP client communication"
	flagClientRPCPort       = "client-rpc-port"
	flagClientRPCPortDesc   = "port used for client RPC client communication"
	flagClusterSerfPort     = "cluster-port"
	flagClusterSerfPortDesc = "port used for internal nodes communication"
	flagClusterRPCPort      = "cluster-rpc-port"
	flagClusterRPCPortDesc  = "port used for internal RPC communication"
	flagClusterRaftPort     = "cluster-raft-port"
	flagClusterRaftPortDesc = "port used for internal RAFT communication"
	flagDataDir             = "data-dir"
	flagDataDirDesc         = "data directory"
	flagDebug               = "debug"
	flagDebugShort          = "D"
	flagDebugDesc           = "enable debug logging"
	flagDebugRaft           = "debug-raft"
	flagDebugRaftDesc       = "enable raft protocol debug logging"
	flagDebugSerf           = "debug-serf"
	flagDebugSerfDesc       = "enable serf protocol debug logging"
	flagDirector            = "director"
	flagDirectorDesc        = "enable director server on this node"
	flagEncrypt             = "encrypt"
	flagEncryptDesc         = "set the gossip encryption key"
	flagFile                = "file"
	flagFileShort           = "f"
	flagFileDesc            = "source file"
	flagID                  = "id"
	flagIDDesc              = "identifier of the resource"
	flagJoin                = "join"
	flagJoinDesc            = "address of a node to join a cluster. Can be specified as a comma separated list"
	flagLogFormat           = "log-format"
	flagLogFormatDesc       = "set the log format [console, json]"
	flagNoGraph             = "no-graph"
	flagNoGraphDesc         = "disable graph server on this node"
	flagServerAddr          = "addr"
	flagServerAddrDesc      = "server address"
	flagVerbose             = "verbose"
	flagVerboseDesc         = "enable verbose logging"
)

const (
	defaultBindAddr        = "0.0.0.0"
	defaultBootstrapExpect = 1
	defaultClusterPort     = 44011
	defaultClusterRaftPort = 44012
	defaultClusterRPCPort  = 44013
	defaultDataDir         = "~/.grog"
	defaultClientHTTPPort  = 8000
	defaultClientRPCPort   = 9000
	logFormatConsole       = "console"
	logFormatJSON          = "json"
)

var defaultServerAddr = fmt.Sprintf("127.0.0.1:%d", defaultClientRPCPort)
