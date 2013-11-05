package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	conf_lazyquots := flag.Bool("lazy_quotes", false, "Parse csv with lazy quote rules")
	conf_trim_leading_spaces := flag.Bool("trim_leading_space", false, "Trim leading spaces")
	conf_columns := flag.Int("columns", -1, "Set number of columns (-1 disables all checks")
	flag.Parse()

	p := csv.NewReader(os.Stdin)
	p.LazyQuotes = *conf_lazyquots
	p.TrimLeadingSpace = *conf_trim_leading_spaces
	p.FieldsPerRecord = *conf_columns
	lineno := 1
	for {
		line, err := p.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%d -", lineno)
		for num, elem := range line {
			fmt.Printf(" %d:%q", num, elem)
		}
		fmt.Println()
		lineno++
	}
}
