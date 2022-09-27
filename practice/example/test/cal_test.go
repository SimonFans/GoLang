package cal

import "testing"

func TestAddNumber(t *testing.T) {
	// call cal.go
	res := addNumber(5)
	if res != 10 {
		t.Fatalf("AddNumber() result wrong. Expectation=%v, Actual=%v\n", 10, res)
	}
	// if no problem, show logs
	t.Logf("AddNumber() implement correct")
}

func TestMinusNumber(t *testing.T) {
	// call cal.go
	res := minusNumber(12, 10)
	if res != 2 {
		t.Fatalf("AddNumber() result wrong. Expectation=%v, Actual=%v\n", 2, res)
	}
	// if no problem, show logs
	t.Logf("AddNumber() implement correct")
}
