package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCopyFile(t *testing.T) {
	// Create a temporary source file with some content
	sourceContent := []byte("This is the content of the source file.")
	sourceFile, err := ioutil.TempFile("", "source.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary source file: %v", err)
	}
	defer os.Remove(sourceFile.Name())
	defer sourceFile.Close()
	if _, err := sourceFile.Write(sourceContent); err != nil {
		t.Fatalf("Failed to write to source file: %v", err)
	}

	// Create a temporary destination file
	destinationFile, err := ioutil.TempFile("", "destination.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary destination file: %v", err)
	}
	defer os.Remove(destinationFile.Name())
	defer destinationFile.Close()

	// Call the function to copy the file
	err = copyFile(sourceFile.Name(), destinationFile.Name())
	if err != nil {
		t.Fatalf("Error while copying file: %v", err)
	}

	// Read content of destination file and compare it with the original content
	destinationContent, err := ioutil.ReadFile(destinationFile.Name())
	if err != nil {
		t.Fatalf("Error reading destination file: %v", err)
	}
	if string(destinationContent) != string(sourceContent) {
		t.Fatalf("Content of source and destination files do not match")
	}
}

func TestCopyFileNonexistentSource(t *testing.T) {
	// Call the function with a nonexistent source file
	err := copyFile("nonexistent.txt", "destination.txt")
	if err == nil {
		t.Fatalf("Expected an error but got none")
	}
}

func TestCopyFileNoPermission(t *testing.T) {
	// Create a source file with read-only permission
	sourceFile, err := ioutil.TempFile("", "source.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary source file: %v", err)
	}
	defer os.Remove(sourceFile.Name())
	defer sourceFile.Close()
	if err := sourceFile.Chmod(0400); err != nil {
		t.Fatalf("Failed to set permissions on source file: %v", err)
	}

	// Call the function with the read-only source file
	err = copyFile(sourceFile.Name(), "destination.txt")
	if err != nil {
		t.Fatalf("Got an error but expected none")
	}
}

func TestMoveFile(t *testing.T) {
	// Create a temporary source file with some content
	sourceContent := []byte("This is the content of the source file.")
	sourceFile, err := ioutil.TempFile("", "source.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary source file: %v", err)
	}
	defer os.Remove(sourceFile.Name())
	defer sourceFile.Close()
	if _, err := sourceFile.Write(sourceContent); err != nil {
		t.Fatalf("Failed to write to source file: %v", err)
	}

	// Specify the destination file path
	destinationFilePath := "./destination.txt"

	// Call the function to move the file
	err = moveFile(sourceFile.Name(), destinationFilePath)
	if err != nil {
		t.Fatalf("Error while moving file: %v", err)
	}

	// Check if source file exists
	if _, err := os.Stat(sourceFile.Name()); !os.IsNotExist(err) {
		t.Fatalf("Source file still exists after moving: %s", sourceFile.Name())
	}

	// Check if destination file exists
	if _, err := os.Stat(destinationFilePath); os.IsNotExist(err) {
		t.Fatalf("Destination file doesn't exist after moving: %s", destinationFilePath)
	}

	// Read content of destination file and compare it with the original content
	destinationContent, err := ioutil.ReadFile(destinationFilePath)
	if err != nil {
		t.Fatalf("Error reading destination file: %v", err)
	}
	if string(destinationContent) != string(sourceContent) {
		t.Fatalf("Content of source and destination files do not match")
	}
}

func TestMoveFileNonexistentSource(t *testing.T) {
	// Call the function with a nonexistent source file
	err := moveFile("nonexistent.txt", "destination.txt")
	if err == nil {
		t.Fatalf("Expected an error but got none")
	}
}
