package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Your solution goes here. Good luck!
	filesNames := listFiles("testdata")

	for _, fileName := range filesNames {
		fmt.Println(fileName)
	}
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
