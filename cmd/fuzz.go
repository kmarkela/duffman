package cmd

import (
	"log"

	"github.com/kmarkela/duffman/internal/fuzz"
	"github.com/kmarkela/duffman/internal/pcollection"
	"github.com/spf13/cobra"
)

// fuzzCmd represents the fuzz command
var fuzzCmd = &cobra.Command{
	Use:   "fuzz",
	Short: "fuzz all endpoint from Postman Collection",
	Long:  `It allows to fuzz muptiple parameters over multiple endpoints`,
	Run: func(cmd *cobra.Command, args []string) {

		// get params
		collF, err := cmd.Flags().GetString("collection")
		if err != nil {
			log.Fatalln(err)
		}
		envF, err := cmd.Flags().GetString("enviroment")
		if err != nil {
			log.Fatalln(err)
		}

		fname, err := cmd.Flags().GetString("wordlist")
		if err != nil {
			log.Fatalln(err)
		}

		greet()

		coll, err := pcollection.New(collF, envF)
		if err != nil {
			log.Fatalln(err)
		}

		proxy, err := cmd.Flags().GetString("proxy")
		if err != nil {
			log.Fatalln(err)
		}
		workers, err := cmd.Flags().GetInt("workers")
		if err != nil {
			log.Fatalln(err)
		}
		maxReq, err := cmd.Flags().GetInt("maxReq")
		if err != nil {
			log.Fatalln(err)
		}
		headers, err := cmd.Flags().GetStringSlice("headers")
		if err != nil {
			log.Fatalln(err)
		}
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			log.Fatalln(err)
		}

		f, err := fuzz.New(workers, maxReq, headers, proxy, verbose)
		if err != nil {
			log.Fatalln(err)
		}

		f.Run(&coll, fname)

	},
}

func init() {

	fuzzCmd.Flags().StringP("wordlist", "l", "", "wordlits to fuzz")
	fuzzCmd.MarkFlagRequired("wordlist")
	fuzzCmd.Flags().StringP("proxy", "p", "", "proxy")
	fuzzCmd.Flags().IntP("maxReq", "m", 0, "max amount of requests per second")
	fuzzCmd.Flags().IntP("workers", "w", 10, "amount of workers")
	fuzzCmd.Flags().StringSlice("headers", []string{}, "replace header if exists, add if it wasn't in original request")

	rootCmd.AddCommand(fuzzCmd)
}
