package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SendRequest() {
	client := http.Client{}
	msg := strings.NewReader("Simon")
	url := "http://localhost:8082/cs"
	req, err := http.NewRequest("PUT", url, msg)
	if err != nil {
		fmt.Println("Error: creating request!!!", err)
		return
	}
	req.Header.Add("Content-Type", "text/plain")
	// client send request now
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: sending request!!!", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: reading response body!!!", err)
		return
	}
	fmt.Println("Response from Server: ", string(b))
}

func main() {
	SendRequest()
}
