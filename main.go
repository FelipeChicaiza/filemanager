package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func createFile(filename, content string) {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	fmt.Println("File created successfully:", filename)
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Content of", filename, ":\n", string(data))
}

func deleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted successfully:", filename)
}

func listFiles() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Error listing files:", err)
		return
	}
	fmt.Println("Files in current directory:")
	for _, file := range files {
		fmt.Println(" -", file.Name())
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [create|read|delete|list] [filename] [content]")
		return
	}

	command := os.Args[1]

	switch command {
	case "create":
		if len(os.Args) < 4 {
			fmt.Println("Usage: create <filename> <content>")
			return
		}
		createFile(os.Args[2], os.Args[3])

	case "read":
		if len(os.Args) < 3 {
			fmt.Println("Usage: Read <filename>")
			return
		}
		readFile(os.Args[2])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: Delete <filename>")
			return
		}
		deleteFile(os.Args[2])

	case "list":
		listFiles()
	default:
		fmt.Println("Unknown command:", command)
	}
}
