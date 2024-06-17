package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "v0.1.0"

var rootCmd = &cobra.Command{
	Use:   "DuffMan",
	Short: "A simple CLI application",
	Long:  `This is a simple CLI application with Cobra.`,
	// This function will run if no subcommands are provided
	Run: func(cmd *cobra.Command, args []string) {
		Greet()
		fmt.Println("This is the default command.")
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "V", false, "Verbose")
	rootCmd.PersistentFlags().StringP("collectionFile", "f", "", "path to collection file")
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	// Get Cli Parameters
	// var flags = &CliFlags{}
	// flags.Parse()

	// print version and exit
	// if flags.Version {
	// 	fmt.Printf("DuffMan - version %s", version)
	// 	return
	// }

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
