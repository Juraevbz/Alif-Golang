package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	if age < 7 {
		fmt.Println("Free")
	} else if age >= 7 && age < 18 {
		fmt.Println("50% discound")
	} else if age > 64 {
		fmt.Println("25% discound")
	} else {
		fmt.Println("Full price")
	}
}
