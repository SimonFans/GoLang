package Basic

import (
	"fmt"
	"log"
	"time"
)

func timeit(name string) func() {
	start := time.Now()
	fmt.Println("starting count time...")
	return func() {
		duration := time.Since(start)
		log.Printf("%s took %s ", name, duration)
	}
}

func Dot(v1, v2 []float64) (float64, error) {
	defer timeit("dot")()

	if len(v1) != len(v2) {
		return 0, fmt.Errorf("dot of different size vectors")
	}

	d := 0.0
	for i, val1 := range v1 {
		val2 := v2[i]
		d += val1 * val2
	}
	fmt.Println("For end ----")

	return d, nil
}
