package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5: ")
	ratingInput := bufio.NewReader(os.Stdin)
	rating, _ := ratingInput.ReadString('\n')
	fmt.Println("Thank you for rating our pizza: ", rating)

	// Convert string to float
	ratingAsFloat, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("Added one to your rating: ", ratingAsFloat+1)
	}

}

// TrimSpace is used to remove any leading or trailing whitespace from the input.
