package main

import "fmt"

func main() {
	var p1 = f()
	var p2 = f()
	fmt.Printf("f(): %v; f(): %v; f() == f()? %t\n", p1, p2, p1 == p2)
}

func f() *int {
	v := 1
	return &v
}
