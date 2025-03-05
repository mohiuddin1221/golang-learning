package main

import (
	"fmt"
)

const a = 10
var p = 30

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}
	add(5, 10)
	add(p, a)
}

func init() {
	fmt.Println("hello")
}

// main function
func main() {
	call()
	fmt.Println(a)
}
