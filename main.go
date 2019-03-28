package main

import (
	"bufio"
	"flag"
	"fmt"
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
	questionOne(names, firstnames, lastnames)
	questionTwo(firstnames)
}

func questionOne(n, f, l se.List) {
	fmt.Printf("\n\t*** Question #1 ***\n")
	fmt.Printf("There are %d unique full names\n", n.GetUniqueElementsCount())
	fmt.Printf("There are %d unique first names\n", f.GetUniqueElementsCount())
	fmt.Printf("There are %d unique last names\n", l.GetUniqueElementsCount())
}

func questionTwo(f se.List) {
	fmt.Printf("\n\t*** Question #2 ***\n")
	fmt.Printf("The ten most common first names are:\n")
	names := f.GetTopElements(10)
	for i := 0; i < 10; i++ {
		fmt.Printf("\t%s (%d)\n", names[i].Name, names[i].Weight)
	}
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
