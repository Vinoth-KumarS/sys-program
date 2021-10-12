package main

import (
	"fmt"
	"unicode"
)

func main() {
	var a byte
	fmt.Println("Are you sure? Do you want to close the terminal! (y/n)")
	fmt.Scanf("%c", &a)
	a = byte(unicode.ToLower(rune(a)))
	switch a {
	case 'y':
		fmt.Println("Terminal closed successfully!")
	case 'n':
		fmt.Println("Terminal returns to work!")
	default:
		fmt.Println("Invalid Character..")
	}
}
