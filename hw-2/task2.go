package main

import "fmt"

func main() {
	var hourOfDay int
	fmt.Print("Enter the current hour (0-23): ")
	fmt.Scan(&hourOfDay)

	if hourOfDay >= 5 && hourOfDay < 12 {
		fmt.Println("It's morning")
	} else if hourOfDay >= 12 && hourOfDay < 17 {
		fmt.Println("It's afternoon")
	} else if hourOfDay >= 17 && hourOfDay < 21 {
		fmt.Println("It's evening")
	} else if (hourOfDay >= 21 && hourOfDay <= 23) || (hourOfDay >= 0 && hourOfDay < 5) {
		fmt.Println("It's night")
	} else {
		fmt.Println("Invalid hour")
	}
}

