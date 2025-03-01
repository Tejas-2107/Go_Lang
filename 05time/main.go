package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Current time: ")
	currTime :=time.Now();
	fmt.Println(currTime.Format("02-01-2006 15:04:05 Monday"));
}