package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "transaction-manager",
	Short: "transaction-manager manage application runtime",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "conf", "", "", "config file path")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Printf("failed to execute command. err: %v", err)
		os.Exit(1)
	}
}
