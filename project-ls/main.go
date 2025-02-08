package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Your solution goes here. Good luck!
	directory := "testdata"
	files := listFiles(directory)
	list := strings.Join(files, " ")

	fmt.Println(list)
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
