package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Your solution goes here. Good luck!
	http.HandleFunc("/hello", GreetHandler)

	log.Println("Listen at port :8080")
	err := http.ListenAndServe(":8080", nil)

	log.Fatal(err)
}

func GreetHandler(writer http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprint(writer, "Hello, ", name)
}
