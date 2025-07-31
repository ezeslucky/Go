package main

import "fmt"

// for loop example
func main() {
	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// infinite loop example (commented out to prevent unreachable code error)
	// for {
	// 	fmt.Println("This will run forever")
	// }

	//for loop 
	for j := 0; j < 5; j++ {
		// break example
		if j == 3 {
			break // exits the loop when j is 3
		}
		fmt.Println(j)
		// continue example
		if j == 1 {
			continue // skips the rest of the loop when j is 1
		}
		fmt.Println("After continue:", j)
	}



	// range loop example
	numbers := []int{1, 2, 3, 4, 5}
	for i, v := range numbers {
		fmt.Println("Index:", i, "Value:", v)
	}
}
	