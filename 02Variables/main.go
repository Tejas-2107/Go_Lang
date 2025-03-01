package main

import (
	"fmt"
)
//public
const LoggedInToken string = "xyz"  
func main() {
	var name string = "tejas"
	fmt.Println(name)
	fmt.Printf("varible of type: %T \n", name)

	var length, _ = fmt.Println(name)
	fmt.Println(length)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("varible of type: %T \n", isLoggedIn)

	//default values and aliases

	var defaultInt float32
	fmt.Println(defaultInt)

	//implicit type

	var myName = "tejas"
	fmt.Println(myName)
	//mynName = 3 gives an error

	//no var style
	// := this operator is called walrus operator
	// it is used to declare and assign value to a variable
	totalUsers := 2000
	fmt.Println(totalUsers)

	fmt.Printf("%b \n", 42424242)

}
