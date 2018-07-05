package bootstrap

import (
	"github.com/spf13/cobra"
	"fmt"
)


func init()  {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		RunE:  runVersion,
	}
	rootCmd.AddCommand(versionCmd)
}

func runVersion(_ *cobra.Command, _ []string) error {
	fmt.Printf("grog - v%s [build=%s]\n", appVersion, appBuild)
	return nil
}
