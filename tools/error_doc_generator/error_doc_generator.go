//go:build error_doc_generator

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func findFilesWithName(startPath, fileName string) ([]string, error) {
	var foundFiles []string

	err := filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == fileName {
			foundFiles = append(foundFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return foundFiles, nil
}

func main() {

	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}
	targetPath := filepath.Join(currentPath, "..", "..", "x")

	// get all errors.go in x folder
	errorFile := "errors.go"
	filePaths, _ := findFilesWithName(targetPath, errorFile)
	if len(filePaths) == 0 {
		fmt.Println("Not find target files in x folder")
		os.Exit(1)
	}

	// get module name and bind with paths (one module may have multiple errors.go)
	moduleWithPaths := make(map[string][]string)
	for _, filePath := range filePaths {
		moduleName := findModuleName(filePath)
		if moduleName == "" {
			fmt.Printf("Failed to get module name for %s\n", filePath)
			os.Exit(1)
		}
		moduleWithPaths[moduleName] = append(moduleWithPaths[moduleName], filePath)
	}

	filePath := targetPath + "/errors.md"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// category
	file.WriteString("<!-- TOC -->\n")
	file.WriteString("Category\n")
	columnTemplate := "  * [%s](#%s)\n"
	for moduleName := range moduleWithPaths {
		file.WriteString(fmt.Sprintf(columnTemplate, strings.Title(moduleName), moduleName))
	}
	file.WriteString("<!-- TOC -->\n")

	// errors in each module
	for moduleName, filePaths := range moduleWithPaths {

		// table header
		file.WriteString("\n")
		file.WriteString("## " + strings.Title(moduleName) + "\n")
		file.WriteString("\n")
		file.WriteString("|Error Name|Codespace|Code|Description|\n")
		file.WriteString("|:-|:-|:-|:-|\n")

		for _, filePath := range filePaths {
			errDict := getErrors(filePath)
			moduleName := getModuleNameValue(filePath)

			for _, errInfo := range errDict {
				file.WriteString(errInfo.toString(moduleName))
			}
		}
	}
}
