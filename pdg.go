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
		if info.IsDir() {
			// Skip the top-level directory itself
			if relPath != "." {
				treeOutput.WriteString(fmt.Sprintf("%s├── %s/\n", indent, info.Name()))
			}
		} else {
			treeOutput.WriteString(fmt.Sprintf("%s├── %s\n", indent, info.Name()))
		}
		return nil
	})

	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	return treeOutput.String()
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run pdir-generator.go <top_level_directory> <root_directory> <exclude_directory1,exclude_directory2,...>")
		os.Exit(1)
	}

	topLevelDirectory := os.Args[1]
	rootDirectory := os.Args[2]
	directoriesToExclude := strings.Split(os.Args[3], ",")

	// Generate the directory structure
	directoryStructure := generateDirectoryStructure(rootDirectory, directoriesToExclude)

	// Print the top-level directory and the directory structure
	fmt.Println(topLevelDirectory + "/")
	fmt.Print(directoryStructure)
}

/*

GENERATE THE PROJECT DIRECTORY LOCATION(IOS AND LINUX):

pwd

COMMAND EXAMPLE:

go run pdg.go classroom-management /Users/tavito/Documents/go/classroom-management .git,pdg.go

*/
