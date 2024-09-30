package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	if age > 0 && age <= 6 {
		fmt.Println("Preschooler")
	} else if age >= 7 && age <= 17 {
		fmt.Println("School-age child")
	} else if age >= 18 && age <= 24 {
		fmt.Println("Student")
	} else if age >= 25 && age <= 64 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior")
	}
}