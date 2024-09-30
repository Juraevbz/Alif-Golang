package main

import "fmt"

func main() {
	var side1, side2, side3 float64
	fmt.Print("Enter the lengths of the three sides of the triangle: ")
	fmt.Scan(&side1, &side2, &side3)

	if (side1+side2 > side3) && (side1+side3 >side2) && (side2+side3 > side1) {
		fmt.Println("The triangle exists")
	} else {
		fmt.Println("The triangle does not exists")
	}
}