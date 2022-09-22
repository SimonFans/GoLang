package Basic

import (
	"fmt"
	"io"
	"log"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
}

// genReport generates a fixed with report in the format
// Symbol: 10 chars, left padded
// Volume: 4 digits, 0 padded
// Price: 2 digits after the decimal

func GenReport(w io.Writer, trades []Trade) {
	for i, t := range trades {
		log.Printf("%d: %#v", i, t)
		// ... 2: main.Trade{Symbol:"BRK-A", Volume:1, Price:399100}
		fmt.Fprintf(w, "%-10s %04d %.2f\n", t.Symbol, t.Volume, t.Price)
		// MSFT       0231 234.57
	}
}
