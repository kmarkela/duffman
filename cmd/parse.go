package cmd

import (
	"fmt"
	"log"

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

		coll, err := pcollection.CollFromJson(collF, envF)
		if err != nil {
			log.Fatalln(err)
		}

		// temp
		fmt.Println("Vars:")
		// fmt.Println(coll)
		for _, v := range coll.Variables {
			fmt.Printf("%s: %s\n", v.Key, v.Value)
		}
		fmt.Println("Requests:")
		for _, v := range coll.Requests {
			fmt.Printf("Method: %s\n URL: %s\n, body: %v\n, headers: %v\n", v.Method, v.URL.Raw, v.Body, v.Header)
			fmt.Println("===========================")
		}

	},
}

func init() {
	parseCmd.Flags().Bool("br", false, "brief")
	rootCmd.AddCommand(parseCmd)
}

// func printBrif() {

// }
