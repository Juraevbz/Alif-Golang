package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	x, y = y, x
	fmt.Println(x, y)
}