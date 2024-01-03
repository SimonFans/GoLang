package main

import (
	"fmt"
	"io"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	info, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("Server failed to parse client's info\n", err)
		return
	}
	switch r.Method {
	case http.MethodPut:
		if r.URL.Path == "/" {
			fmt.Fprintf(w, "Welcome to Homepage: %s\n", string(info))
		} else if r.URL.Path == "/cs" {
			fmt.Fprintf(w, "Welcome to CS Program: %s\n", string(info))
		} else {
			fmt.Fprintf(w, "sorry %s, this is an invalid URL\n", string(info))
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	myHandler := &MyHandler{}
	server := http.Server{
		Addr:    "localhost:8082",
		Handler: myHandler,
	}
	fmt.Println("Server is listening on :8082")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server Listen Error: ", err)
	}
}
