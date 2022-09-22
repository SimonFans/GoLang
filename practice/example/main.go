package main

import (
	"example/Method"
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("Hello GO!!")
	// fmt.Println(Basic.KillContainer("data/containerId"))
	// // write data to csv format
	// output_name := "items.csv"
	// items := []Basic.Item{
	// 	{"m183x", "Magic Wand"},
	// 	{"m184y", "Invisibility Cape"},
	// 	{"m185z", "Levitation Spell"},
	// }
	// fmt.Print(Basic.WriteItems(output_name, items))
	// nums := []int{1}
	//fmt.Println(secondToLast(nums)) // will panic
	// fmt.Println(safeSecondToLast(nums))
	// find out odd number in an array
	// values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// fmt.Println(Basic.Filter(Basic.IsOdd, values))
	// Time functions
	// date := time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC)
	// fmt.Println(date, "&", date.Weekday()) // 2021-12-31 00:00:00 +0000 UTC Friday
	// fmt.Println("The type of time.Saturday =>", reflect.TypeOf(time.Saturday))
	// nbd := Basic.NextBusinessDay(date)
	// fmt.Println(nbd, nbd.Weekday())
	// defer example testing
	//testDefer()
	// measure.go
	// v := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(Basic.Dot(v, v))
	// time conversion based on location
	// out, err := Basic.TsConvert("2021-03-08T19:12", "America/Los_Angeles", "Asia/Jerusalem")
	// if err != nil {
	// 	fmt.Printf("error: %s", err)
	// 	return
	// }
	// fmt.Println(out)
	// report.go
	// trades := []Basic.Trade{
	// 	{"MSFT", 231, 234.57},
	// 	{"TSLA", 123, 686.75},
	// 	{"BRK-A", 1, 399100},
	// }
	// fmt.Println(trades)
	// Basic.GenReport(os.Stdout, trades)
	// Get length of characters rather than length of list
	// words := []string{"«", "Don't", "Panic", "»"}
	// fmt.Println(Basic.GetLineLength(words))
	// greek.go
	// fmt.Println(Basic.EnglishFor("p"))
	// regular expression
	// line := "12 shares of MSFT for $234.57"
	// t, err := Basic.ParseLine(line)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", t) // {Symbol:MSFT Volume:12 Price:234.57}
	// grep.go
	// file, err := os.Open("data/journal.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// matches, err := Basic.Grep(file, "System is rebooting")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%d reboot\n", len(matches))
	// history.go
	// freqs, err := Basic.CmdFreq("zsh_history")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for cmd, cnt := range freqs {
	// 	fmt.Printf("%s -> %d\n", cmd, cnt)
	// }
	// // events.go
	// evt, err := Method.NewDoorEvent("front door", time.Now(), "open")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // &{Event:{ID:front door Time:2021-04-30 14:47:40.31038 +0300 IDT m=+0.000170354} Action:open}
	// fmt.Printf("%+v\n", evt)
	// sensor.go
	// t := Method.Thermostat{"Living Room", 18.2}
	// fmt.Printf("%s before: %.2f\n", t.ID, t.Value)
	// t.SetValue(20.12)
	// fmt.Printf("%s after: %.2f\n", t.ID, t.Value)
	// interface.go
	// circle := Method.Circle{0, 0, 5}
	// rectangle := Method.Rectangle{5, 10}
	// shape := []Method.Shape{&circle, &rectangle}
	// Method.PrintAll(shape)
	//counter.go
	// Method.RecordEvent(&Method.ClickEvent{})
	// Method.RecordEvent(&Method.ClickEvent{})
	// Method.RecordEvent(&Method.HoverEvent{})
	// Method.RecordEvent(3)
	// fmt.Println("Event counts:", Method.EventCounts)
	// json/serialize.go
	// monster := Json.Monster{"Niu Mo Wang", 18}
	// fmt.Println(monster)
	// jsonStr, err := json.Marshal(monster)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(jsonStr))
	// Method study
	// var p1 Method.Person
	// p1.Name = "Simon"
	// p1.Speak() // transmit param just like copy
	// p1.Jisuan()
	// p1.Jisuan2(10)
	// fmt.Printf("Two sum result: %v\n", p1.GetSum(2, 8))
	// var c Method.Circle
	// c.Radius = 4.0
	// res := c.Area() // equivalent to (&c).Area()
	// fmt.Println("the Circle area is", res)
	// var val Method.ChangeValue = 10
	// fmt.Printf("val: %v\n", val)
	// val.Change()
	// fmt.Println("After change, the new value is", val)
	// // Since you defined string() method, it will call that by default using the format you defined
	// student1 := Method.Student{Name: "Ximeng", Age: 31}
	// fmt.Println(&student1)
	// // determine if odd or even
	//mu := Method.Methodutil{}
	// fmt.Println((&mu).IsOdd(7))
	// 9*9 math table
	//mu.PrintTable(9)
	// Print chars
	// (&mu).PrintChars(7, 2, "@")
	// // Calculator
	// calculator := Method.Calcualtor{Num1: 2.4, Num2: 7.6}
	// operator := "-"
	// fmt.Printf("The operator is %s, result=%.2f\n", operator, calculator.GetOperator(operator))
	// homework
	//var arr = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//mu.Transpose(arr)
	// factory mode. To deal with any struct name, or field name starts with small case letter -> teacher
	var t = Method.NewTeacher("Ximeng", 5.0)
	fmt.Printf("Name=%s, Rate=%.2f", t.Name, t.GetRate())
	var a = []int{3, 4, 1}
	sort.Ints(a)
	fmt.Println(a)

}

type Person struct {
	Age     int
	Name    string
	Courses []int
	ptr     *int
	map1    map[string]string
	slice   []int
}

func test(test_name string) {
	fmt.Println("test name is", test_name)
}

func testDefer() {
	defer test("test")
	a := 3
	b := 4
	fmt.Println("a plus b is", a+b)

}

func safeSecondToLast(nums []int) (i int, err error) {
	defer func() {
		if e := recover(); e != nil { // e is interface{}
			err = fmt.Errorf("%v", e)
		}
	}()

	return secondToLast(nums), nil
}

func secondToLast(nums []int) int {
	return nums[len(nums)-2]
}
