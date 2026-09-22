package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	allLines := getFile("sillyRec.txt")
	for i, line := range allLines {
		if i%2 != 0 {
			fmt.Println(line)
		}
	}
}

func getFile(name string) []string {
	contents, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	bigString := string(contents)
	allLines := strings.Split(bigString, "\n")
	return allLines
}
