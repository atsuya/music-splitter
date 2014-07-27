package main

import (
	"flag"
	"os"
)

func main() {
	// parse commandline options
	music := flag.String("music", "", "a path to music")
	csv := flag.String("csv", "", "a path to csv")
	flag.Parse()

	if *music == "" && *csv == "" {
		flag.Usage()
		os.Exit(0)
	}
}
