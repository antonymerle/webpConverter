package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Check if the directory path was provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a directory path as an argument.")
		return
	}

	// Get the directory path from the command-line arguments
	dir := os.Args[1]

	// Create the "convertedWebp" subdirectory if it doesn't exist
	convertedDir := filepath.Join(dir, "convertedWebp")
	if _, err := os.Stat(convertedDir); os.IsNotExist(err) {
		os.Mkdir(convertedDir, 0755)
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is a WebP file
		if strings.HasSuffix(path, ".webp") {
			fmt.Println("Converting", path)

			// Build the command to execute ffmpeg
			outputPath := strings.TrimSuffix(path, ".webp") + ".jpg"
			cmd := exec.Command("ffmpeg", "-i", path, outputPath)

			// Run the command
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error:", err)
			}

			// Move the original WebP file to the "convertedWebp" subdirectory
			newPath := filepath.Join(convertedDir, filepath.Base(path))
			err = os.Rename(path, newPath)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
