package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/transaction-manager/internal/config"
	"github.com/transaction-manager/internal/deamon"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:          "server",
	Short:        "Start a http server.",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load(configFile)
		if err != nil {
			log.Fatal(err)
		}

		deamon.RunHttpServer(cfg)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
