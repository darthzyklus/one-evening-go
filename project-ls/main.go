package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Your solution goes here. Good luck!
	filesNames := listFiles("testdata")
	fmt.Println(strings.Join(filesNames, " "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, _ := ioutil.ReadDir(dirname)

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
