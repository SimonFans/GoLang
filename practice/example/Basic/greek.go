package Basic

import (
	"fmt"
	"strings"
)

type Letter struct {
	Symbol  string
	English string
}

var letters = []Letter{
	{"Σ", "Sigma"},
}

func EnglishFor(greek string) (string, error) {
	for _, letter := range letters {
		if strings.EqualFold(greek, letter.Symbol) {
			return letter.English, nil
		}
	}
	return "", fmt.Errorf("Unknown greek letter: %#v", greek)
}
