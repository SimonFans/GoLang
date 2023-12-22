package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard logger") // 2023/12/22 17:28:51 standard logger
	// log.Printf("%s", "Print format")
	// log.Fatal("exit after this line")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with mrio seconds") // 2023/12/22 17:32:14.206623 with mrio seconds
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")
	mylog := log.New(os.Stdout, "ximeng logger:", log.LstdFlags)
	mylog.Println("from mylog")
	mylog.SetPrefix("ximeng logger2:")
	mylog.Println("updating prefix")
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello")
	buf.Truncate(28) // from buflog: buf:2023/12/22 18:15:20 hell
	fmt.Println("from buflog:", buf.String())
	jsonHandler := slog.NewJSONHandler(os.Stdin, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	myslog.Info("hello again", "key", "val", "age", 25)
}
