package main

import (
	"flag"
)

func main() {
	copyFlag := flag.Bool("copy", false, "Copy files/directories")
	moveFlag := flag.Bool("move", false, "Move files/directories")
	listFlag := flag.Bool("list", false, "List files/directories")
	deleteFlag := flag.Bool("delete", false, "Delete files")
	searchFlag := flag.String("search", "", "Search for files/directories")

	flag.Usage = help

	flag.Parse()

	// Perform operations based on flags
	if *copyFlag {
		copy()
	} else if *moveFlag {
		move()
	} else if *listFlag {
		listFiles(flag.Args())
	} else if *searchFlag != "" {
		searchFiles(flag.Args(), *searchFlag)
	} else if *deleteFlag {
		delete()
	} else {
		help()
	}
}
