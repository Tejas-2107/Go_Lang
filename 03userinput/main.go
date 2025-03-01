package main

import (
	"bufio"
	"fmt"
	"os"
)

func main1() {
	fmt.Println("Taking input from user")

	// Create a buffered reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name: ")
	// ReadString will read until a newline character is found
	//comma ok or comma error syntax 
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is: ", name)
	fmt.Printf("Type of name is %T \n", name)


}
