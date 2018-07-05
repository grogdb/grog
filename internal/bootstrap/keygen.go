package bootstrap

import (
	"github.com/spf13/cobra"
	"github.com/grogdb/grog/internal/util/random"
	"encoding/base64"
	"fmt"
)

func init()  {
	var keygenCmd = &cobra.Command{
		Use:   "keygen",
		Short: "Generate a new encryption key",
		RunE:  runKeygen,
	}
	rootCmd.AddCommand(keygenCmd)
}

func runKeygen(_ *cobra.Command, _ []string) error {
	key := base64.StdEncoding.EncodeToString(random.Bytes(16))
	fmt.Println(key)
	return nil
}