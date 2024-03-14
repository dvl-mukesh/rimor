# Rimor: File System Exploration Command Line Tool

Rimor is a command-line tool written in Go for exploring the file system. The name "Rimor" comes from Latin, meaning explorer. With Rimor, you can navigate through directories, view file contents, search for files, and perform various file system operations, all from the comfort of your terminal.

## Features

- **Navigate**: Traverse through directories with simple commands.
- **View**: View the contents of files directly in the terminal.
- **Search**: Search for files based on name or content.
- **Manipulate**: Perform basic file system operations like copy, move, delete, etc.

## Installation

To install Rimor, you'll need to have Go installed on your system. Then, simply run:

```bash
go get -u github.com/dvl-mukesh/rimor
```

## Usage

```bash
./rimor -copy sourcefile destinationfile
./rimor -move sourcefile destinationfile
./rimor -list directory
./rimor -search directory searchquery
./rimor -delete filename
```
