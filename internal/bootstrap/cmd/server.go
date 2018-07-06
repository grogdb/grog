package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start in server mode",
		RunE:  runServer,
	}
	serverCmd.Flags().String(flagAdvertiseAddr, "", flagAdvertiseAddrDesc)
	serverCmd.Flags().String(flagBindAddr, defaultBindAddr, flagBindAddrDesc)
	serverCmd.Flags().Int(flagClusterPort, defaultClusterPort, flagClusterPortDesc)
	serverCmd.Flags().String(flagDataDir, defaultDataDir, flagDataDirDesc)
	serverCmd.Flags().String(flagEncrypt, "", flagEncryptDesc)
	serverCmd.Flags().Int(flagGossipPort, defaultGossipPort, flagGossipPortDesc)
	serverCmd.Flags().String(flagJoin, "", flagJoinDesc)

	rootCmd.AddCommand(serverCmd)
}

func runServer(_ *cobra.Command, _ []string) error {
	// FIXME: Missing implementation
	return nil
}
