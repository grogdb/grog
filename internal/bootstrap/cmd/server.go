package cmd

import (
	"fmt"
	"net"
	"strings"

	"github.com/grogdb/grogdb/internal/bootstrap"
	"github.com/hashicorp/go-sockaddr/template"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start server",
		RunE:  runServer,
	}

	serverCmd.Flags().String(flagAdvertiseAddr, "", flagAdvertiseAddrDesc)
	serverCmd.Flags().String(flagBindAddr, defaultBindAddr, flagBindAddrDesc)
	serverCmd.Flags().Int(flagBootstrapExpect, defaultBootstrapExpect, flagBootstrapExpectDesc)
	serverCmd.Flags().Int(flagClientHTTPPort, defaultClientHTTPPort, flagClientHTTPPortDesc)
	serverCmd.Flags().Int(flagClientRPCPort, defaultClientRPCPort, flagClientRPCPortDesc)
	serverCmd.Flags().Int(flagClusterSerfPort, defaultClusterPort, flagClusterSerfPortDesc)
	serverCmd.Flags().Int(flagClusterRPCPort, defaultClusterRPCPort, flagClusterRPCPortDesc)
	serverCmd.Flags().Int(flagClusterRaftPort, defaultClusterRaftPort, flagClusterRaftPortDesc)
	serverCmd.Flags().String(flagDataDir, defaultDataDir, flagDataDirDesc)
	serverCmd.Flags().Bool(flagDebugRaft, false, flagDebugRaftDesc)
	serverCmd.Flags().Bool(flagDebugSerf, false, flagDebugSerfDesc)
	serverCmd.Flags().Bool(flagDirector, false, flagDirectorDesc)
	serverCmd.Flags().Bool(flagNoGraph, false, flagNoGraphDesc)
	serverCmd.Flags().String(flagEncrypt, "", flagEncryptDesc)
	serverCmd.Flags().String(flagJoin, "", flagJoinDesc)

	serverCmd.Flags().MarkHidden(flagDebugRaft)
	serverCmd.Flags().MarkHidden(flagDebugSerf)

	rootCmd.AddCommand(serverCmd)
}

func runServer(_ *cobra.Command, _ []string) error {
	// Advertise address
	advertiseAddr := viper.GetString(flagAdvertiseAddr)
	if advertiseAddr == "" {
		resolvedAdvertiseAddr, err := template.Parse("{{ GetPrivateIP }}")
		if err != nil {
			return fmt.Errorf("failed to resolve advertise IP address: %v", err)
		}
		advertiseAddr = resolvedAdvertiseAddr
	}

	advertiseIPAddr := net.ParseIP(advertiseAddr)
	if advertiseIPAddr == nil {
		return fmt.Errorf("invalid --advertise-addr: %s", advertiseAddr)
	}

	// Bind address
	bindAddr := viper.GetString(flagBindAddr)
	bindIPAddr := net.ParseIP(bindAddr)
	if bindIPAddr == nil {
		return fmt.Errorf("invalid --bind-addr: %s", bindAddr)
	}

	// Client HTTP port
	clientHTTPPort := viper.GetInt(flagClientHTTPPort)
	if !isValidPort(clientHTTPPort) {
		return fmt.Errorf("invalid --client-http-port: %d", clientHTTPPort)
	}

	// Client RPC port
	clientRPCPort := viper.GetInt(flagClientRPCPort)
	if !isValidPort(clientRPCPort) {
		return fmt.Errorf("invalid --client-rpc-port: %d", clientRPCPort)
	}

	// Cluster Serf port
	clusterSerfPort := viper.GetInt(flagClusterSerfPort)
	if !isValidPort(clusterSerfPort) {
		return fmt.Errorf("invalid --cluster-port: %d", clusterSerfPort)
	}

	// Cluster RPC port
	clusterRPCPort := viper.GetInt(flagClusterRPCPort)
	if !isValidPort(clusterRPCPort) {
		return fmt.Errorf("invalid --cluster-rpc-port: %d", clientRPCPort)
	}

	// Cluster Raft port
	clusterRaftPort := viper.GetInt(flagClusterRaftPort)
	if !isValidPort(clusterRaftPort) {
		return fmt.Errorf("invalid --cluster-raft-port: %d", clusterRaftPort)
	}

	// Data dir
	dataDir := viper.GetString(flagDataDir)
	if strings.Contains(dataDir, "~") {
		resolvedDataDir, err := homedir.Expand(dataDir)
		if err != nil {
			return fmt.Errorf("failed to expand data directory: %s", dataDir)
		}
		dataDir = resolvedDataDir
	}

	// Join cluster
	joinAddrs := make([]string, 0)
	joinArg := viper.GetString(flagJoin)
	for _, addr := range strings.Split(joinArg, ",") {
		if strings.Contains(addr, ":") {
			joinAddrs = append(joinAddrs, strings.TrimSpace(addr))
		} else {
			joinAddrs = append(joinAddrs, fmt.Sprintf("%s:%d", strings.TrimSpace(addr), clusterSerfPort))
		}
	}

	enableDirector := viper.GetBool(flagDirector)
	disableGraphServer := viper.GetBool(flagNoGraph)

	if !enableDirector && disableGraphServer {
		return fmt.Errorf("neither director or graph services enabled for this node")
	}

	var graphHTTPAddr *net.TCPAddr
	var graphRPCAddr *net.TCPAddr

	if !disableGraphServer {
		graphHTTPAddr = &net.TCPAddr{IP: bindIPAddr, Port: clientHTTPPort}
		graphRPCAddr = &net.TCPAddr{IP: bindIPAddr, Port: clientRPCPort}
	}

	var raftAddr *net.TCPAddr
	var raftAdvertiseAddr *net.TCPAddr

	if enableDirector {
		raftAddr = &net.TCPAddr{IP: bindIPAddr, Port: clusterRaftPort}
		raftAdvertiseAddr = &net.TCPAddr{IP: advertiseIPAddr, Port: clusterRaftPort}
	}

	cfg := &bootstrap.ServerConfig{
		BootstrapExpect:   viper.GetInt(flagBootstrapExpect),
		ClusterRPCAddr:    &net.TCPAddr{IP: bindIPAddr, Port: clusterRPCPort},
		DataDir:           dataDir,
		DirectorEnabled:   enableDirector,
		GraphEnabled:      !disableGraphServer,
		GraphHTTPAddr:     graphHTTPAddr,
		GraphRPCAddr:      graphRPCAddr,
		JoinAddrs:         joinAddrs,
		SerfAddr:          &net.TCPAddr{IP: bindIPAddr, Port: clusterSerfPort},
		SerfAdvertiseAddr: &net.TCPAddr{IP: advertiseIPAddr, Port: clusterSerfPort},
		SerfEncryptionKey: viper.GetString(flagEncrypt),
		SerfDebug:         viper.GetBool(flagDebugSerf),
		RaftAddr:          raftAddr,
		RaftAdvertiseAddr: raftAdvertiseAddr,
		RaftDebug:         viper.GetBool(flagDebugRaft),
	}
	return bootstrap.NewServer(rootLogger, cfg).Start()
}
