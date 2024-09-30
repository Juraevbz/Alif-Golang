package main

import "fmt"

func main() {
	var year int
	fmt.Print("Enter the year: ")
	fmt.Scan(&year)

	if (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0) {
		fmt.Println(year, "is a leap year")
	} else {
		fmt.Println(year, "is not a leap year")
	}
}
