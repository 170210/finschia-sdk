//go:build error_doc_generator

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type errorInfo struct {
	errorName   string
	codeSpace   string
	code        string
	description string
}

func (err errorInfo) toString(moduleName string) string {
	errorInfoTemplate := "|%s|%s|%s|%s|\n"
	if err.errorName == "ModuleName" {
		if moduleName != "" {
			return fmt.Sprintf(errorInfoTemplate, err.errorName, moduleName, err.code, err.description)
		} else {
			fmt.Println("failed to find moduleName")
			os.Exit(1)
		}
	}

	return fmt.Sprintf(errorInfoTemplate, err.errorName, err.codeSpace, err.code, err.description)

}

func getConst(line string) (string, string) {
	line = strings.Replace(line, "const", "", 1)
	parts := strings.Split(line, "=")
	if len(parts) == 2 {
		i := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(strings.Trim(parts[1], `"`))
		return i, val
	} else {
		fmt.Printf("failed to get the value in: %s \n", line)
		os.Exit(1)
	}
	return "", ""
}

func addError(line string, errorDict map[string]string) errorInfo {
	parts := strings.SplitN(line, "=", 2)
	errName := strings.TrimSpace(parts[0])
	errBody := strings.TrimSpace(parts[1])
	// error info is like as sdkerrors.Register(...)
	pattern := regexp.MustCompile(`sdkerrors\.Register\((.*)\)`)
	match := pattern.FindStringSubmatch(errBody)

	if len(match) == 2 {
		parts := strings.SplitN(match[1], ",", 3)

		if len(parts) == 3 {
			codeSpace := strings.TrimSpace(parts[0])
			code := strings.TrimSpace(parts[1])
			description := strings.TrimSpace(parts[2])

			if constValue, found := errorDict[codeSpace]; found {
				codeSpace = constValue
			}

			return errorInfo{
				errorName:   errName,
				codeSpace:   codeSpace,
				code:        code,
				description: description,
			}
		} else {
			fmt.Printf("failed to get error info in: %s \n", line)
			os.Exit(1)
		}
	} else {
		fmt.Printf("failed to parse error info in: %s \n", line)
		os.Exit(1)
	}

	return errorInfo{}
}

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

func findModuleName(s string) string {
	startIndex := strings.Index(s, "/x/") + len("/x/")
	endIndex := strings.Index(s[startIndex:], "/")

	if startIndex != -1 && endIndex != -1 {
		return s[startIndex : startIndex+endIndex]
	}
	return ""
}

func getModuleNameValue(filePath string) string {
	possibleFileNames := []string{"keys.go", "key.go"}
	var keyFilePath string
	for _, fileName := range possibleFileNames {
		paramPath := strings.Replace(filePath, "errors.go", fileName, 1)
		if _, err := os.Stat(paramPath); err == nil {
			keyFilePath = paramPath
			break
		}
	}

	if keyFilePath != "" {
		file, err := os.Open(keyFilePath)
		if err != nil {
			fmt.Printf("%s cannot be opened\n", keyFilePath)
			os.Exit(1)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			// get module name
			if strings.Contains(line, "ModuleName = ") {
				_, val := getConst(line)
				return val
			}
		}
	}

	return ""
}

func getErrors(p string) []errorInfo {
	var errorDict []errorInfo
	constDict := make(map[string]string)

	file, err := os.Open(p)
	if err != nil {
		fmt.Printf("%s cannot be opened\n", p)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// get const
		if strings.Contains(line, "=") {
			if !strings.Contains(line, "sdkerrors.Register") {
				identifier, value := getConst(line)
				constDict[identifier] = value
			} else {
				errorDict = append(errorDict, addError(line, constDict))
			}
		}
	}
	return errorDict
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
