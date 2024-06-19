package cmd

import (
	"log"

	"github.com/kmarkela/duffman/internal/output"
	"github.com/kmarkela/duffman/internal/pcollection"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parse only collection file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		collF, err := cmd.Flags().GetString("collectionFile")
		if err != nil {
			log.Fatalln(err)
		}
		envF, err := cmd.Flags().GetString("enviromentFile")
		if err != nil {
			log.Fatalln(err)
		}

		greet()

		coll, err := pcollection.New(collF, envF)
		if err != nil {
			log.Fatalln(err)
		}

		output.PrintCol(output.Brief, &coll)

	},
}

func init() {
	parseCmd.Flags().Bool("br", false, "brief")
	rootCmd.AddCommand(parseCmd)
}
