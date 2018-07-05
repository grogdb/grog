package bootstrap

import (
	"github.com/spf13/cobra"
)


func init()  {
	var clusterCmd = &cobra.Command{
		Use:   "cluster",
		Short: "Manage the cluster",
		RunE:  runCluster,
	}
	clusterCmd.Flags().String(flagServerAddr, defaultServerAddr, flagServerAddrDesc)
	rootCmd.AddCommand(clusterCmd)
}

func runCluster(_ *cobra.Command, _ []string) error {
	return nil
}
