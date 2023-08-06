package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func testFileHandler() {
	// Get the directory path from the environment variable
	dir := os.Getenv("DIRECTORY_PATH")
	if dir == "" {
		fmt.Println("Directory path not provided.")
		return
	}

	// Perform file operations
	filename := "example.txt"

	// Get file
	_, err := get_file(filepath.Join(dir, filename))
	if err != nil {
		fmt.Printf("Failed to get file: %s\n", err)
	}

	// Create file
	err = create_file(filepath.Join(dir, filename), []byte("Hello, World!"))
	if err != nil {
		fmt.Printf("Failed to create file: %s\n", err)
	}

	// Overwrite file
	err = overwrite_file(filepath.Join(dir, filename), []byte("New content!"))
	if err != nil {
		fmt.Printf("Failed to overwrite file: %s\n", err)
	}

	// Rename file
	newFilename := "new_example.txt"
	err = rename_file(filepath.Join(dir, filename), filepath.Join(dir, newFilename))
	if err != nil {
		fmt.Printf("Failed to rename file: %s\n", err)
	}

	// Delete file
	err = delete_file(filepath.Join(dir, newFilename))
	if err != nil {
		fmt.Printf("Failed to delete file: %s\n", err)
	}
}

// Get file content
func get_file(filepath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// Create file
func create_file(filepath string, content []byte) error {
	err := ioutil.WriteFile(filepath, content, 0666)
	if err != nil {
		return err
	}

	fmt.Println("File created successfully.")
	return nil
}

// Overwrite file content
func overwrite_file(filepath string, content []byte) error {
	err := ioutil.WriteFile(filepath, content, 0666)
	if err != nil {
		return err
	}

	fmt.Println("File overwritten successfully.")
	return nil
}

// Rename file
func rename_file(oldPath, newPath string) error {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}

	fmt.Println("File renamed successfully.")
	return nil
}

// Delete file
func delete_file(filepath string) error {
	err := os.Remove(filepath)
	if err != nil {
		return err
	}

	fmt.Println("File deleted successfully.")
	return nil
}
