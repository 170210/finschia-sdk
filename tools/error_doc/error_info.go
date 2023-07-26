package main

import (
	"bufio"
	"fmt"
	"os"
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
	if err.codeSpace == "ModuleName" {
		if moduleName != "" {
			return fmt.Sprintf(errorInfoTemplate, err.errorName, moduleName, err.code, err.description)
		} else {
			fmt.Println("failed to find moduleName")
			os.Exit(1)
		}
	}

	return fmt.Sprintf(errorInfoTemplate, err.errorName, err.codeSpace, err.code, err.description)

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
			description := strings.Trim(strings.TrimSpace(parts[2]), `"`)

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
