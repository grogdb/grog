package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	var clusterCmd = &cobra.Command{
		Use:   "cluster",
		Short: "Manage the cluster",
	}
	clusterCmd.Flags().String(flagServerAddr, defaultServerAddr, flagServerAddrDesc)
	rootCmd.AddCommand(clusterCmd)

	var clusterMemberListCmd = &cobra.Command{
		Use:   "members",
		Short: "List all cluster members",
		RunE:  runClusterMemberList,
	}
	clusterCmd.AddCommand(clusterMemberListCmd)

	var clusterStatusCmd = &cobra.Command{
		Use:   "status",
		Short: "Get the cluster status",
		RunE:  runClusterStatus,
	}
	clusterCmd.AddCommand(clusterStatusCmd)
}

func runClusterMemberList(_ *cobra.Command, _ []string) error {
	// FIXME: Missing implementation
	return nil
}

func runClusterStatus(_ *cobra.Command, _ []string) error {
	// FIXME: Missing implementation
	return nil
}
