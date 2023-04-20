package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		fmt.Printf("Status: %v\nStatusCode: %v\nBody: %v\n", resp.Status, resp.StatusCode, resp.Body)
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading: %s; %v\n", url, err)
		}
		fmt.Println(string(b))
	}
}

// If success, go run main.go http://www.google.com
Status: 200 OK
StatusCode: 200
Body: &{[] 0xc00020a0c0 <nil> <nil>}

// If fail, go run main.go http://www.googe.com
fetch: Get "https://luj.proasdf.com/?s=poet&d=www.googe.com": x509: “*.proasdf.com” certificate is expired
exit status 1

// another way to use io.copy() 
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error reading URL:", err)
			return
		}
		defer resp.Body.Close()
		fmt.Printf("Status: %v\nStatusCode: %v\nBody: %v\n", resp.Status, resp.StatusCode, resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Println("Error coping response body", err)
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error response status code:", resp.StatusCode)
		}
	}
}
