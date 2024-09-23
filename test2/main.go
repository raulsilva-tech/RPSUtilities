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
	directory := "my_directory" // Change this to your target directory

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

//---------------------------------------

// package main

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	// Define the directory name
// 	dirName := "my_directory"

// 	// Create the directory
// 	err := os.Mkdir(dirName, 0755)
// 	if err != nil {
// 		if os.IsExist(err) {
// 			fmt.Printf("Directory already exists: %s\n", dirName)
// 		} else {
// 			fmt.Printf("Error creating directory: %v\n", err)
// 			return
// 		}
// 	} else {
// 		fmt.Printf("Directory created: %s\n", dirName)
// 	}

// 	// Create 20,000 text files
// 	for i := 20001; i <= 30000; i++ {
// 		fileName := fmt.Sprintf("file_%d.txt", i)
// 		filePath := filepath.Join(dirName, fileName)

// 		// Create and write to the file
// 		err = os.WriteFile(filePath, []byte("This is a sample text file."), 0644)
// 		if err != nil {
// 			fmt.Printf("Error creating file %s: %v\n", fileName, err)
// 			return
// 		}
// 	}

// 	fmt.Println("Successfully created 20,000 text files.")
// }
