package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"regexp"

	se "github.com/kgosse/testforgc/searchengine"
)

var re = regexp.MustCompile(`^([a-zA-Z]+), ([a-zA-Z]+) -- .+\n$`)

func main() {
	f := flag.String("f", "test-input-file.list", "the input file")
	flag.Parse()

	file, err := os.Open(*f)
	defer file.Close()
	handleError(err)

	reader := bufio.NewReader(file)
	names := se.NewList()
	lastnames := se.NewList()
	firstnames := se.NewList()

	var line string
	for {
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		handleError(err)

		if re.MatchString(line) {
			n := re.FindStringSubmatch(line)
			l, f := n[1], n[2]
			names.Add(l + ", " + f)
			lastnames.Add(l)
			firstnames.Add(f)
		}
	}
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
