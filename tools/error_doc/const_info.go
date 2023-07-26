package main

import (
	"fmt"
	"os"
	"strings"
)

func getConst(line string) (string, string) {
	line = strings.Replace(line, "const", "", 1)
	parts := strings.Split(line, "=")
	if len(parts) == 2 {
		i := strings.TrimSpace(parts[0])
		val := strings.Trim(strings.TrimSpace(parts[1]), `"`)
		return i, val
	} else {
		fmt.Printf("failed to get the value in: %s \n", line)
		os.Exit(1)
	}
	return "", ""
}
