package cmd

import (
	"log"

	"github.com/kmarkela/duffman/internal/interactive"
	"github.com/kmarkela/duffman/internal/pcollection"
	"github.com/spf13/cobra"
)

var interCmd = &cobra.Command{
	Use:   "interactive",
	Short: "execute separate request with manual customisation",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		collF, err := cmd.Flags().GetString("collection")
		if err != nil {
			log.Fatalln(err)
		}

		envF, err := cmd.Flags().GetString("environment")
		if err != nil {
			log.Fatalln(err)
		}

		proxy, err := cmd.Flags().GetString("proxy")
		if err != nil {
			log.Fatalln(err)
		}

		coll, err := pcollection.New(collF, envF, make([]string, 0))
		if err != nil {
			log.Fatalln(err)
		}

		i, err := interactive.New(proxy)
		if err != nil {
			log.Fatalln(err)
		}

		i.Run(&coll)

	},
}

func init() {
	fuzzCmd.Flags().StringP("proxy", "p", "", "proxy")
	rootCmd.AddCommand(interCmd)
}
