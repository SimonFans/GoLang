// This code will send HTTP request and save the response to a file
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// create a new HTTP client (struct)
	client := http.Client{}
	// NewRequest will build *Request(struct)包装request
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)
	fmt.Println(req) // &{GET https://jsonplaceholder.typicode.com/posts HTTP/1.1 1 1 map[] <nil> <nil> 0 [] false jsonplaceholder.typicode.com map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000240b8}
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}
	// do send HTTP request and return a HTTP response(struct){statuscode, body}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("The StatusCode is:%v\n", resp.StatusCode)
	// read the response body
	file, err := os.Create("../os/response.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error coping data from HTTP response...", err)
		return
	}
	// print the response body
	fmt.Println("HTTP response data has been written to the file in os path..")
}
