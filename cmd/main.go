package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:  "transaction-manager",
	Short: "transaction-manager manage application runtime",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "conf", "", "", "config file path")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("failed to execute command. err: %v", err)
		os.Exit(1)
	}
}
