package cmd

import "flag"

type arrayFlags []string

func (i *arrayFlags) String() string {
	//implementation of flag.Value
	return ""
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type CliFlags struct {
	ColFile          string
	Headers          arrayFlags
	Version, Verbose bool
}

func (cf *CliFlags) Parse() {
	flag.StringVar(&cf.ColFile, "f", "", "path to Postman collection")
	flag.Var(&cf.Headers, "h", "Headers")
	flag.BoolVar(&cf.Version, "V", false, "Print Version")
	flag.BoolVar(&cf.Version, "vv", false, "Verbose")
	flag.Parse()
}
