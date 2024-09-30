package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	fmt.Print("Enter the third number: ")
	fmt.Scan(&num3)

	if num1 > num2 && num1 > num3 {
		fmt.Println("The largest number is:", num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Println("The largest number is:", num2)
	} else {
		fmt.Println("The largest number is:", num3)
	}
}