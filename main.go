package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("This assignment belongs to Group E")

	fmt.Println(" Calculate Area of Ractangle ")
	fmt.Println(calculateareaofractangle())

	fmt.Println(" Check Palindrome ")
	palindrome()

}

// Function Three
// created by Veerpal Kaur(500226407)
// this function will Calculate the area of rectangle by entering length and width.
func calculateareaofractangle() float64 {

	var length, width float64

	fmt.Println("Enter the length of the rectangle: ")
	fmt.Scan(&length)
	fmt.Println("Enter the Width of the rectangle: ")
	fmt.Scan(&width)

	return length * width
}

// Function Seven
// created by Gurcharan Singh(500228108)
// this function will tells weather the entered string is palidrome or not by using strings
func palindrome() {
	var word string

	fmt.Print("Enter the string without space: ")
	fmt.Scan(&word)

	word = strings.ToLower(word)

	n := ""

	for i := len(word) - 1; i >= 0; i-- {
		n = n + string(word[i])
	}

	if n == word {
		fmt.Println("String is Palindrome")
	} else {
		fmt.Println("String is not Palindrome")
	}
}
