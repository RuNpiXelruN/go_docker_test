package main

import (
	"fmt"
	"io"
	"net/http"
)

var outText *string

func init() {
	// const (
	// 	defaultText = "Hello World!"
	// 	usage       = "enter custom text for the program to print"
	// )
	// outText = flag.String("text", defaultText, usage)
	// outText = flag.String("t", defaultText, usage+" (shorthand)")
	// flag.Parse()
}

func main() {
	fmt.Println("Running on port 3333...")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3333", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!")
}
