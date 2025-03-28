package main
import(
	"fmt"
)
//pass by value
//pass by refference
func print(numbers *[5]string){
	fmt.Println(numbers)

}
type User struct{
	name string
	age int
	email string
	address string
}

func main(){
	instace := User{
		name: "mohiuddin",
        age:  30,
        email: "mohiuddin@gmail.com",
        address: "dhaka",
	}

	p := &instace
	fmt.Println(p.name)
	fmt.Println(instace)
	// pointer
	// pointer is address of memory(ram or harddisk)
	// x := 20
	// fmt.Println(x)
    // // address of  == &(ampersand)
	// address := &x
	// *address = 50
	// fmt.Println("address:",address)
	// fmt.Println("value at address:",*address)

	arr := [5]string{"1", "2", "3", "4", "5"}
	print(&arr)


}