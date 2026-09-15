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
	nextClass := classInfo{}
	cursor := bufio.NewReader(os.Stdin)
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
	//fmt.Print("Enter another class name: ")
	//name, err = cursor.ReadString('\n')
	//if err != nil {
	//	panic(err)
	//}
	//classes = append(classes, name)
	//fmt.Print("Enter another class name: ")
	//name, err = cursor.ReadString('\n')
	//if err != nil {
	//	panic(err)
	//}
	//classes = append(classes, name)
	fmt.Println(classes)
}
