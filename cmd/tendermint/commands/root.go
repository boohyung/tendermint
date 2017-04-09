package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tendermint/go-logger"
	tmcfg "github.com/tendermint/tendermint/config/tendermint"
)

var config *viper.Viper
var log = logger.New("module", "main")

//global flag
var logLevel string

var RootCmd = &cobra.Command{
	Use:   "tendermint",
	Short: "Tendermint Core (BFT Consensus) in Go",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Set("log_level", logLevel)
	},
}

func init() {

	// setup configuration
	config = tmcfg.GetConfig("")

	//parse flag and set config
	RootCmd.PersistentFlags().StringVar(&logLevel, "log_level", config.GetString("log_level"), "Log level")

	// set the log level
	logger.SetLogLevel(config.GetString("log_level"))
}
