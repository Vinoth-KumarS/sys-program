package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int

	fmt.Println("Enter the Number: ")
	fmt.Scanf("%d", &num)
	fmt.Println(isDivided(num))

}

func isDivided(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "The number is divided by 3 and 5"
	} else if n%3 == 0 {
		return "The number is divided by 3 only"
	} else if n%5 == 0 {
		return "The number is divided by 5 only"
	}
	return strconv.Itoa(n)
}
