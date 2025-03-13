package main

import (
	"fmt"
)
type User struct{
	Name string
	Age  int
}
func PrintUser(usr User){
	fmt.Println("name :", usr.Name)
	fmt.Println("Age :", usr.Age)
}

func (usr User) PrintUserReciver(){
	fmt.Println("name :", usr.Name)
	fmt.Println("Age :", usr.Age)
}

func main(){
	var user1 User

	user1 = User{
		Name: "Topu",
		Age:  40,
	}

	user1.PrintUserReciver()

}