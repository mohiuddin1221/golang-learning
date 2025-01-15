package main

import (
	"fmt"
)

func WelcomeMessage() {
	fmt.Println("Welcome to the Application")
}

func GetUserName() string {
	var name string
	fmt.Println("Please Enter your Name:")
	fmt.Scanln(&name)
	return name
}

func GetTwoNumbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter Your First Number:")
	fmt.Scanln(&num1)
	fmt.Println("Enter Your Second Number:")
	fmt.Scanln(&num2)
	return num1, num2
}

func SumNumbers(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func Display(name string, sum int) {
	fmt.Println("Hello", name)
	fmt.Println("The sum of numbers is:", sum)
}

func GoodBye() {
	fmt.Println("Thank you for using the Application")
	fmt.Println("GoodBye!")
}

func main() {
	WelcomeMessage()
	name := GetUserName()
	num1, num2 := GetTwoNumbers()
	sum := SumNumbers(num1, num2)
	Display(name, sum)
	GoodBye()
}
