package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

var url, account, dir string

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(describeaccountbalance)

}