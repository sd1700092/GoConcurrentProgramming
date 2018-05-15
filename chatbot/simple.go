package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		os.Exit(1)
	}
	input = input[:len(input)-1]
	fmt.Printf("Hello %s! What can I do for you?", input)

	for {
		input, err := reader.ReadString('\n')
		if err!=nil{
			fmt.Printf("An error occurred: %s\n", err)
			continue
		}
		input = input[:len(input) - 1]
		input = strings.ToLower(input)
		switch input {
		case "": continue
		case "bye", "nother":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Sorry, I didn't catch you.")
		}
	}
}
