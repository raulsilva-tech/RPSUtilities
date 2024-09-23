package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Specify the directory you want to check
	directory := "//192.168.7.204/videos" // Change this to your target directory

	// Call the function to find the oldest file
	oldestFile, err := findOldestFile(directory)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if oldestFile != "" {
		fmt.Printf("The oldest file is: %s\n", oldestFile)
	} else {
		fmt.Println("No files found in the directory.")
	}
}

// findOldestFile finds the oldest file in the given directory
func findOldestFile(dir string) (string, error) {
	log.Println("0")
	entries, err := os.ReadDir(dir)

	log.Println("1")
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, _ := entry.Info()

		infos = append(infos, info)
	}
	if err != nil {
		return "", err
	}
	log.Println("2")
	var oldestFile string
	var oldestTime time.Time

	// count := 0
	for _, file := range infos {
		// We only consider regular files
		if file.Mode().IsRegular() {
			if oldestFile == "" || file.ModTime().Before(oldestTime) {
				oldestFile = filepath.Join(dir, file.Name())
				oldestTime = file.ModTime()
			}
			// count++
			// fmt.Println(count)
		}

	}
	log.Println(oldestFile)
	return oldestFile, nil
}
