package Method

import (
	"fmt"
)

/*
type Person struct {
	Name string
}

func (p person) Set(){
	fmt.Println(p.Name)
}
(1) func (p Person) Set(){} means struct Person has a method called Set()
(2) (p Person) indicates method Set() bind with type Person
*/

type Person struct {
	Name string
}

func (p Person) Speak() {
	fmt.Println(p.Name, "is learning Go")
}

func (p Person) Jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Printf("%v calculate 1 to 1000, the result is %d\n", p.Name, res)
}

func (p Person) Jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Printf("%v calculate 1 to %v, the result is %d\n", p.Name, n, res)
}

func (p Person) GetSum(n1, n2 int) int {
	return n1 + n2
}

type Circle1 struct {
	radius float64
}

// in order to improve efficiency
func (c *Circle) Area() float64 {
	c.Radius = 10 // c saves the address or point to the address, not value copy
	return 3.14 * c.Radius * c.Radius
}

type ChangeValue int

func (val *ChangeValue) Change() {
	*val = *val + 1
}

// overwrite fmt.println method by defining the string() method
type Student struct {
	Name string
	Age  int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name = %v, Age = %v", stu.Name, stu.Age)
	return str
}

// odd or even
type Methodutil struct {
}

func (mu *Methodutil) IsOdd(num int) bool {
	if num%2 == 1 {
		return true
	} else {
		return false
	}
}

// print specific chatacters in a given row and column
func (mu *Methodutil) PrintChars(row int, col int, key string) {
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

// 9*9 Math table
func (mu *Methodutil) PrintTable(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			fmt.Printf("%vx%v=%v ", j, i, j*i)
			if j == i {
				break
			}
		}
		fmt.Println()
	}
}

func (mu *Methodutil) Transpose(arr [][]int) {
	row, col := len(arr), len(arr[0])
	fmt.Println(row)
	fmt.Println(col)
	for r := 0; r < row/2; r++ {
		for c := r; c < col; c++ {
			temp := arr[c][r]
			arr[c][r] = arr[r][c]
			arr[r][c] = temp
		}
	}
	fmt.Println("After transpose:")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}

}

// calculator do '+', '-', '*' and '/'
type Calcualtor struct {
	Num1 float64
	Num2 float64
}

func (calculator *Calcualtor) GetOperator(operator string) float64 {
	var res float64
	switch operator {
	case "+":
		res = calculator.Num1 + calculator.Num2
	case "-":
		res = calculator.Num1 - calculator.Num2
	case "*":
		res = calculator.Num1 * calculator.Num2
	case "/":
		res = calculator.Num1 / calculator.Num2
	default:
		fmt.Println("operator is wrong")
	}
	return res
}

// factory mode. To deal with any struct name, or field name starts with small case letter
type teacher struct {
	Name string
	rate float64
}

// deal with struct name teacher starts with small case letter t => define a func
func NewTeacher(n string, r float64) *teacher {
	return &teacher{
		Name: n,
		rate: r,
	}
}

// deal with field name rate starts with small case letter r => define a func
func (t *teacher) GetRate() float64 {
	return t.rate
}
