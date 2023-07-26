package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
