package cmd

import (
	"log"

	"github.com/kmarkela/duffman/internal/interactive"
	"github.com/kmarkela/duffman/internal/pcollection"
	"github.com/spf13/cobra"
)

var interCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Lightweight CLI postman client",
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

		coll, err := pcollection.New(collF, envF, nil)
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
	interCmd.Flags().StringP("proxy", "p", "", "proxy")
	rootCmd.AddCommand(interCmd)
}
