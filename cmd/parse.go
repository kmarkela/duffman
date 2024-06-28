package cmd

import (
	"log"
	"strings"

	"github.com/kmarkela/duffman/internal/output"
	"github.com/kmarkela/duffman/internal/pcollection"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parse only collection file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		collF, err := cmd.Flags().GetString("collection")
		if err != nil {
			log.Fatalln(err)
		}
		envF, err := cmd.Flags().GetString("enviroment")
		if err != nil {
			log.Fatalln(err)
		}

		greet()

		coll, err := pcollection.New(collF, envF, nil)
		if err != nil {
			log.Fatalln(err)
		}

		o, _ := cmd.Flags().GetString("output")

		otype := output.Req
		switch strings.ToLower(o) {
		case "brief":
			otype = output.Brief
		case "full":
			otype = output.Full
		}

		output.PrintCol(otype, &coll)

	},
}

func init() {
	parseCmd.Flags().StringP("output", "", "req", "output type. Possible values: brief, req, full")
	rootCmd.AddCommand(parseCmd)
}
