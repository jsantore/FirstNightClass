package main

import (
	"fmt"
	"os"
	"strings"
)

func getFile() []string {
	contents, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	bigString := string(contents)
	allLines := strings.Split(bigString, "\n")
	return allLines
}
