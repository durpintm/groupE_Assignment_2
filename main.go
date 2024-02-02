package main

import "fmt"

func main() {
	fmt.Println("This assignment belongs to Group E")
	fmt.Println(" Calculate Area of Ractangle ")
	fmt.Println(calculateareaofractangle())

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
