package bootstrap

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
	"github.com/rs/zerolog"
	"os"
	"github.com/rs/zerolog/log"
	"fmt"
)

var rootCmd = &cobra.Command{
	Use:               "grog",
	Short:             "Grog is a fast distributed graph database",
	PersistentPreRunE: initialize,
	// RunE:              runClient,
}

var (
	appBuild = "dev"
	appVersion = "X.X.X"
)

func Execute(build, version string) error {
	appBuild = build
	appVersion = version

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func init()  {
	rootCmd.PersistentFlags().BoolP(flagDebug, flagDebugShort, false, flagDebugDesc)
	rootCmd.PersistentFlags().String(flagLogFormat, logFormatConsole, flagLogFormatDesc)
	rootCmd.PersistentFlags().Bool(flagVerbose, false, flagVerboseDesc)
}

func initialize(cmd *cobra.Command, _ []string) error {
	// cli
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	envKeyReplacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(envKeyReplacer)

	viper.SetEnvPrefix("GROG")
	viper.AutomaticEnv()

	// logging
	if viper.GetBool(flagVerbose) {
		viper.Set(flagDebug, true)
	}

	if viper.GetBool(flagDebug) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	logFormat := viper.GetString(flagLogFormat)
	switch logFormat {
	case logFormatConsole:
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	case logFormatJSON:
		log.Logger = log.With().
			Str("app", "grog").
			Str("version", appVersion).
			Str("build", appBuild).
			Logger()

		zerolog.TimeFieldFormat = ""
	default:
		return fmt.Errorf("invalid log format: %s", logFormat)
	}
	rootLogger = &log.Logger

	return nil
}
