package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const version = "v0.2.1-alpha"

var rootCmd = &cobra.Command{
	Use:   "DuffMan",
	Short: "Fuzzer for Postman collections",
	Long:  `Diagnostic Utility for Fuzzing and Fault Management of API Nodes`,
	Run: func(cmd *cobra.Command, args []string) {
		parseCmd.Execute()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "V", false, "Verbose")
	rootCmd.PersistentFlags().StringP("collectionFile", "f", "", "path to collection file")
	rootCmd.MarkFlagRequired("collectionFile")
	rootCmd.PersistentFlags().StringP("enviromentFile", "e", "", "path to enviroment file")
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
