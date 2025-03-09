package main

import (
    "fmt"
)
const a = 100

var p = 80

func outer() func(){
	money : = 40
	age := 30

	show = func(){
		money = money + a + p
		fmt.Println("Money : ", money)
	}
	return show
}

func call(){
	inc1 := outer()
	inc1()
	inc1()

	inc2 := outer()
	inc2()
	inc2()
}
func main() {
	call()
}
func init() {
	fmt.Println("Hello, world!")

}


func main() {

}