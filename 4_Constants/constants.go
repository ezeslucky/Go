package main
import "fmt"
func main() {
	const xyz string = "lucky"
	fmt.Println(xyz)

	const age int = 30
	fmt.Println(age)

	const (
		port = 3000
		host = "localhost"
	)
	fmt.Println("Port:", port)
	fmt.Println("Host:", host)
}