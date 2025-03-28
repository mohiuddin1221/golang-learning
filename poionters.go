package main
import(
	"fmt"
)


func main(){
	// pointer
	// pointer is address of memory(ram or harddisk)
	x := 20
	fmt.Println(x)
    // address of  == &(ampersand)
	address := &x
	*address = 50
	fmt.Println("address:",address)
	fmt.Println("value at address:",*address)


}