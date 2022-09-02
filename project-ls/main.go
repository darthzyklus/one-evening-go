package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Your solution goes here. Good luck!
	config := Config{
		showHiddenFiles: *flag.Bool("a", false, "show hidden files"),
	}

	flag.Parse()

	app := App{
		config: config,
	}

	filesNames := app.listFiles("testdata")

	for _, fileName := range filesNames {
		fmt.Println(fileName)
	}
}

type Config struct {
	showHiddenFiles bool
}

type App struct {
	config Config
}

func (app App) listFiles(dirname string) []string {
	var dirs []string

	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		if strings.HasPrefix(f.Name(), ".") && app.config.showHiddenFiles {
			continue
		}

		dirs = append(dirs, f.Name())

	}

	return dirs
}
