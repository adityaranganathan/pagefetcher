package main

import (
	"flag"
)

func main() {
	var maxRoutines int
	flag.IntVar(&maxRoutines, "parallel", 10, "maximum number of parallel routines")
	flag.Parse()

	urls := flag.Args()
	urls = AddHTTPScheme(urls)

	ProcessURLs(urls, maxRoutines, FetchPage, PrintResult)
}
