package cal

// command to run in the test folder, the order of cal_test.go and cal.go doesn't matter
// go test -v cal_test.go cal.go
// if you only want to test a specific function in a .go file, then use command: go test -v -test.run TestAddNumber
// no need to have main function because testing framework has included the main in it

func addNumber(n int) int {
	total := 0
	for i := 1; i < n; i++ {
		total += i
	}
	return total
}

func minusNumber(n1, n2 int) int {
	return n1 - n2
}
