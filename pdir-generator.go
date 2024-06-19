package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func generateDirectoryStructure(rootDir string, excludeDirs []string) string {
	var treeOutput strings.Builder

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Exclude specified directories
		for _, excludeDir := range excludeDirs {
			if info.IsDir() && info.Name() == excludeDir {
				return filepath.SkipDir
			}
		}

		// Calculate indentation
		relPath, _ := filepath.Rel(rootDir, path)
		indent := strings.Repeat("│   ", strings.Count(relPath, string(filepath.Separator)))

		// Append to treeOutput
		treeOutput.WriteString(fmt.Sprintf("%s├── %s\n", indent, info.Name()))
		return nil
	})

	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	return treeOutput.String()
}

func main() {
	// Specify the root directory of your project
	rootDirectory := "/Users/tavito/Documents/go/classroom-management"

	// List of directories to exclude (customize as needed)
	directoriesToExclude := []string{"node_modules", ".git", "build"}

	// Generate the directory structure (excluding specified directories)
	directoryStructure := generateDirectoryStructure(rootDirectory, directoriesToExclude)

	// Print the modified directory structure
	fmt.Println(directoryStructure)
}
