package main

import (
    "fmt"
)

// Standard function
func add(x int, y int){
	fmt.Println(x + y)
}

func main(){
	add(5, 10)
	// Anonymous function
	// IIFS Function immediately invokes a function expression
	func (x int, y int){
		fmt.Println(x + y)
	}(30, 40)

    
}
// ini function
func init(){
	fmt.Println("First Print invoke function")
	
}