package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	classes := make([]string, 0)
	cursor := bufio.NewReader(os.Stdin)
	fmt.Print("Enter class name: ")
	name, err := cursor.ReadString('\n')
	if err != nil {
		panic(err)
	}
	classes = append(classes, name)
	fmt.Print("Enter another class name: ")
	name, err = cursor.ReadString('\n')
	if err != nil {
		panic(err)
	}
	classes = append(classes, name)
	fmt.Print("Enter another class name: ")
	name, err = cursor.ReadString('\n')
	if err != nil {
		panic(err)
	}
	classes = append(classes, name)
	fmt.Println(classes)
}
