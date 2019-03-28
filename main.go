package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	f := flag.String("f", "test-input-file.list", "the input file")
	flag.Parse()

	file, err := os.Open(*f)
	defer file.Close()
	handleError(err)
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
