package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Your solution goes here. Good luck!
	directory := "testdata"
	files := listFiles(directory)

	for _, file := range files {
		fmt.Println(file)
	}
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
