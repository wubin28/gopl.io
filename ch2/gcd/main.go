package main

import "fmt"

func main() {
	fmt.Printf("gcd(2,3) is %v\n", gcd(2,3))
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
