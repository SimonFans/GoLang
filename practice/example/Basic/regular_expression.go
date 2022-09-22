package Basic

import (
	"fmt"
	"regexp"
	"strconv"
)

type Transaction struct {
	Symbol string
	Volume int
	Price  float64
}

var transRe = regexp.MustCompile(`(\d+) shares of ([A-Z]+) for \$(\d+(\.\d+)?)`)

func ParseLine(line string) (Transaction, error) {
	matches := transRe.FindStringSubmatch(line)
	fmt.Println(matches)
	if matches == nil {
		return Transaction{}, fmt.Errorf("bad line: %q", line)
	}
	var t Transaction
	t.Symbol = matches[2]
	t.Volume, _ = strconv.Atoi(matches[1])
	t.Price, _ = strconv.ParseFloat(matches[3], 64)
	return t, nil
}
