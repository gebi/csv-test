package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	p := csv.NewReader(os.Stdin)
	flag.BoolVar(&p.LazyQuotes, "lazy_quotes", false, "Parse csv with lazy quote rules")
	flag.BoolVar(&p.TrimLeadingSpace, "trim_leading_space", false, "Trim leading spaces")
	flag.IntVar(&p.FieldsPerRecord, "columns", -1, "Set number of columns (-1 disables all checks")
	flag.Parse()

	lineno := 1
	for {
		record, err := p.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%.2d -", lineno)
		for num, elem := range record {
			fmt.Printf(" %d:%q", num, elem)
		}
		fmt.Println()
		lineno++
	}
}
