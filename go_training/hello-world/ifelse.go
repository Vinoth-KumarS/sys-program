package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Printf("Enter Your age: ")
	fmt.Scanf("%d", &age)

	if age > 18 {
		fmt.Println("You are major")
	} else {
		fmt.Println("You are minor")
	}
	fmt.Println("Enter the Name: ")
	fmt.Scanf("%d", &age)

}
