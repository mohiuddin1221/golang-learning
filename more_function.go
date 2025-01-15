package main
import(
	"fmt"
)
func PrintSomething(){
	fmt.Println("Education must be free")
}
func SayHello(name string){
	fmt.Println("Welcome to Golang Course, ", name)
    
}
func GoodBye(){
	fmt.Println("Goodbye")
}

func main() {
	PrintSomething()
	SayHello("Mohiuddin")
	GoodBye()


}