package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Your solution goes here. Good luck!
	filesNames := listFiles("testdata")
	fmt.Println(strings.Join(filesNames, " "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
