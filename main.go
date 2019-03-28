package main

import (
	"flag"
	"fmt"
)

func main() {
	f := flag.String("f", "test-input-file.list", "the input file")
	flag.Parse()
	fmt.Println(*f)
}
