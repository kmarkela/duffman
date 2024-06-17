package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print Version",
	Long:  "Print Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("# DuffMan: Diagnostic Utility for Fuzzing and Fault Management of API Nodes\n[*] Version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
