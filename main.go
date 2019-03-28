package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`^([a-zA-Z]+), ([a-zA-Z]+) -- .+\n$`)

func main() {
	f := flag.String("f", "test-input-file.list", "the input file")
	flag.Parse()

	file, err := os.Open(*f)
	defer file.Close()
	handleError(err)

	reader := bufio.NewReader(file)

	var line string
	cpt := 0
	for {
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		handleError(err)

		if re.MatchString(line) {
			n := re.FindStringSubmatch(line)
			l, f := n[1], n[2]
			fmt.Printf("%s, %s\n", l, f)
			cpt++
		}

		if cpt == 10 {
			break
		}

	}
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
