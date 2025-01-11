package main

import (
	"fmt"
)

func add(num1 int, num2 int) (int, int) {
	x := num1 + num2
	y := num1*num2
	return x, y
}

func main() {
	a := 10
	b := 20

	sum , mul := add(a, b)
	fmt.Printf("The sum is: %d\n", sum)
	fmt.Printf("The sum is: %d\n", mul)
}
