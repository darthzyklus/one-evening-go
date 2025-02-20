package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	directory := "testdata"
	all := flag.Bool("a", false, "Show hidden files")

	flag.Parse()

	files := listFiles(directory, *all)

	for _, file := range files {
		fmt.Println(file)
	}
}

func listFiles(dirname string, all bool) []string {
	var dirs []string
	var pointPrefix = "."

	files, err := os.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !all && strings.HasPrefix(f.Name(), pointPrefix) {
			continue
		}

		dirs = append(dirs, f.Name())
	}

	return dirs
}
