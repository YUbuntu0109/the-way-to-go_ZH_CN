package main

import "fmt"


func main() {
	fmt.Println("hello 04.2\n")
	fmt.Println("1 + 2 + 3 =", Sum(1,2))
}

// defined a method
func Sum(a, b int) int {
	const c int = 3
	return a + b + c
}
