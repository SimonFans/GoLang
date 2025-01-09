package main

import (
	"flag"
	"fmt"
)

var (
	name  string
	age   int
	debug bool
)

func init() {
	flag.StringVar(&name, "name", "defaultName", "enter your name")
	flag.IntVar(&age, "age", 18, "enter your age")
	flag.BoolVar(&debug, "debug", false, "if enable debug")
}

type Phone struct {
	ID     string
	*Usage // even though phone does not implement GetCPU() directly, but Usage implements as embed, so still valid
}

type Usage struct {
	CPU int
}

type Getter interface {
	GetCPU() int
	GetID() string
}

func (p *Phone) GetID() string {
	if p.ID == "" {
		return "123"
	}
	return p.ID
}

func (u *Usage) GetCPU() int {
	return 500
}

func GetPhoneInfo(getter Getter) {
	fmt.Printf("ID: %s\nCPU: %d\n", getter.GetID(), getter.GetCPU())
}

func main() {
	flag.Parse()
	arguments := make(map[string]string)
	flag.Visit(func(f *flag.Flag) {
		arguments[f.Name] = f.Value.String()
	})
	for name, val := range arguments {
		fmt.Println(name, val)
	}
	required := []string{"name", "debug"}
	for _, key := range required {
		if arguments[key] == "" {
			fmt.Println("missing key: ", key)
		}
	}

	newPhone := &Phone{}
	GetPhoneInfo(newPhone)

	// Create a context with a timeout of 5 seconds
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel() // Cancel context to release resources

	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		fmt.Println("Context done:", ctx.Err()) // Prints: "Context done: context deadline exceeded"
	// 		return
	// 	case <-time.After(1 * time.Second):
	// 		fmt.Println("Still working...")
	// 	}
	// }
}
