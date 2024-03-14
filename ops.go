package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dvl-mukesh/rimor/banner"
)

func copy() {
	if len(flag.Args()) != 2 {
		fmt.Println("Usage: go run main.go -copy source destination")
		return
	}
	src := flag.Args()[0]
	dst := flag.Args()[1]

	info, err := os.Stat(src)
	if err != nil {
		fmt.Println("Unable to read source file")
	}
	if info.IsDir() {
		err = copyDir(src, dst)
	} else {
		err = copyFile(src, dst)
	}

	if err != nil {
		fmt.Printf("Error copying file: %v\n", err)
	} else {
		fmt.Printf("Copied %s to %s\n", src, dst)
	}

}

func help() {
	banner.PrintBanner()
	flag.PrintDefaults()

	fmt.Fprintf(os.Stderr, "\nUsage: go run main.go [flags] [files/directories...]")
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

func delete() {
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: go run main.go -delete file")
		return
	}
	file := args[0]
	err := deleteFile(file)
	if err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
	} else {
		fmt.Printf("Deleted file %s\n", file)
	}
}

func deleteFile(file string) error {
	return os.Remove(file)
}

func move() {
	if len(flag.Args()) != 2 {
		fmt.Println("Usage: go run main.go -move source destination")
		return
	}
	source := flag.Args()[0]
	destination := flag.Args()[1]
	err := moveFile(source, destination)
	if err != nil {
		fmt.Printf("Error moving file: %v\n", err)
	} else {
		fmt.Printf("Moved %s to %s\n", source, destination)
	}
}

func moveFile(src, dst string) error {
	err := copyFile(src, dst)
	if err != nil {
		return err
	}
	return os.Remove(src)
}

func listFiles(paths []string) {
	for _, path := range paths {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Printf("Error listing %s: %v\n", path, err)
			continue
		}
		fmt.Printf("Contents of %s:\n", path)
		for _, file := range files {
			fmt.Printf("\t%s\n", file.Name())
		}
	}
}

func searchFiles(paths []string, searchQuery string) {
	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Contains(info.Name(), searchQuery) {
				fmt.Println(path)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error searching %s: %v\n", path, err)
		}
	}
}

func copyDir(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}
