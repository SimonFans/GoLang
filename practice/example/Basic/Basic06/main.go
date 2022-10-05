package main

import (
	"fmt"
)

// for loop
func main() {
	i := 1
	upper := 100
	cnt := 0
	var slice []int
	for ; i <= upper; i++ {
		if i%9 == 0 {
			slice = append(slice, i)
			cnt++
		}
	}
	fmt.Printf("numbers can be divided by 9:%v, how many: %v\n", slice, cnt)
	// random generate 1-5, count how many times it can reach to 5
	// cnt1 := 0
	// for {
	// 	rand.Seed(time.Now().Unix())
	// 	n := rand.Intn(5) + 1
	// 	cnt1++
	// 	if n == 5 {
	// 		break
	// 	}
	// }
	// fmt.Printf("%v\n times it reach to 5", cnt1)
	// use lable to tell break which for loop you want to break
lable2:
	for i := 0; i < 4; i++ {
		fmt.Printf("i=%v\n", i)
		for j := 0; j < 10; j++ {
			if j == 3 {
				break lable2
			}
			fmt.Printf("j=%v\n", j)
		}
	}
	// goto example: goto <label>
	var n int = 20
	fmt.Println("20")
	if n > 10 {
		goto label1
	}
	fmt.Println("ok1")
label1:
	fmt.Println("label1")
	// return. If return is in the main function, then quit the main. If return is in a regular function, then quit that function
	for i := 1; i < 10; i++ {
		if i == 3 {
			return
		}
		fmt.Println(i)
	}
	fmt.Println("Main ends")
}
