package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// get string length
	str1 := "hello"
	fmt.Println("string length=", len(str1))

	// print each character in the string
	str2 := "hello"
	r := []rune(str2) // [104 101 108 108 111 119 111 114 108 100]
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i])
	}

	// string to integer
	n, err := strconv.Atoi("12")
	if err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println("success converting string to integer", n)
	}

	// integer to string
	str3 := strconv.Itoa(123)
	fmt.Printf("str=%v, type=%T\n", str3, str3)

	// string to bytes
	bytes := []byte("hellogo")
	fmt.Printf("bytes=%v\n", bytes)

	// []bytes to string
	str4 := string([]byte{98, 99, 100})
	fmt.Printf("str=%v\n", str4)

	// decimal to binary, Oct, Hex
	str5 := strconv.FormatInt(123, 2)
	fmt.Printf("123 to bianry is=%v\n", str5)
	str6 := strconv.FormatInt(123, 16)
	fmt.Printf("123 to hex is=%v\n", str6)

	// determine if the substring is in a given string
	fmt.Println(strings.Contains("seafood", "foo"))

	// count # of given chars in a string
	num := strings.Count("aaapple", "a")
	fmt.Printf("there are %v a in the current string\n", num)

	// compare two strings if they are equal to each other
	fmt.Println(strings.EqualFold("abc", "Abc")) // not case sensitive
	fmt.Println("abc" == "Abc")                  // case sensitive

	// return the index of the first appearnace of a char in a string
	index := strings.Index("abcd", "bc")
	fmt.Printf("The first index is=%v\n", index)

	// return the index of the last appearnace of a char in a string. If cannot find it, return -1
	last_index := strings.LastIndex("abcdd", "d")
	fmt.Printf("The last index is=%v\n", last_index)

	// replace one substring to another substring. 1: replace 1 ge. -1: replace all.
	str7 := "go go Olympics"
	str8 := strings.Replace(str7, "go", "Bejing", 1)
	fmt.Printf("old string=%v, new string=%v\n", str7, str8)

	// split string to string array
	strArr := strings.Split("hello,world,abc", ",")
	fmt.Printf("strArr=%v\n", strArr)
	for i, val := range strArr {
		fmt.Printf("index=%v, value=%v\n", i, val)
	}

	// string to lowecase and uppercase
	fmt.Println(strings.ToLower("helloGO"))
	fmt.Println(strings.ToUpper("helloGO"))

	// trim space two sides
	str9 := strings.TrimSpace("  Hi Michael  ")
	fmt.Printf("str=%q\n", str9)

	// trim specific chars on both sides. In this example, remove ! & space on both sides
	str10 := strings.Trim("!he!llo! ", "! ")
	fmt.Printf("str=%q\n", str10)

	// trim single side
	left := strings.TrimLeft("!Hello!", "!")
	right := strings.TrimRight("!Hello!", "!")
	fmt.Printf("left=%q\n", left)
	fmt.Printf("right=%q\n", right)

	// prefix and suffix
	fmt.Println(strings.HasPrefix("ftp://123.56", "ftp"))
	fmt.Println(strings.HasPrefix("ftp://123.56", "htp"))
	fmt.Println(strings.HasSuffix("ftp://123.56", "56"))
	fmt.Println(strings.HasSuffix("ftp://123.56", "456"))
}
