package main

import (
	"fmt"
	"log"
	"net/http"
)

var calls []string
var stats = make(map[string]int)

func main() {
	// Your solution goes here. Good luck!
	http.HandleFunc("/hello", GreetHandler)

	log.Println("Listen at port :8080")
	err := http.ListenAndServe(":8080", nil)

	log.Fatal(err)
}

func GreetHandler(writer http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	calls = append(calls, name)
	stats[name] += 1

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n\n", stats)
	fmt.Fprint(writer, "Hello, ", name)
}
