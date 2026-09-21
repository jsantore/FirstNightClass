package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type classInfo struct {
	Name       string
	Instructor string
	numCredits int
}

func main() {
	classes := make([]classInfo, 0)

	cursor := bufio.NewReader(os.Stdin)
	for {
		nextClass := classInfo{}

		fmt.Print("Enter class name: ")
		name, err := cursor.ReadString('\n')
		if err != nil {
			panic(err)
		}
		nextClass.Name = name
		fmt.Print("Enter instructor name: ")
		teacher, err := cursor.ReadString('\n')
		nextClass.Instructor = teacher
		fmt.Print("Enter number of credits: ")
		creditsAsStr, err := cursor.ReadString('\n')
		creditsAsStr = strings.Trim(creditsAsStr, "\n")
		credits, err := strconv.Atoi(creditsAsStr)
		fmt.Println(err)
		nextClass.numCredits = credits
		classes = append(classes, nextClass)
		fmt.Print("Do you want to enter another class? (y/n): ")
		answer, err := cursor.ReadString('\n')
		if err != nil {
			panic(err)
		}
		answer = strings.Trim(answer, "\n")
		if answer != "y" {
			break
		}
	}

	for i, class := range classes {
		fmt.Printf("Class %d: %s, %s, %d\n", i+1, class.Name, class.Instructor, class.numCredits)
	}
}
