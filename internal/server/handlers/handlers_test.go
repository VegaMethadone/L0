package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestFileContents(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %v\n", err)
	}

	filePath := filepath.Join(wd, "..", "static", "index.html")
	fmt.Println("File path:", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v\n", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatalf("Failed to get file info: %v\n", err)
	}

	fileSize := fileInfo.Size()
	content := make([]byte, fileSize)

	_, err = file.Read(content)
	if err != nil {
		t.Fatalf("Failed to read file: %v\n", err)
	}

	fmt.Println("File contents:")
	fmt.Println(string(content))
}
