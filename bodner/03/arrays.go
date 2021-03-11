package main

import "fmt"

func main() {
	var x [3]int // array of 3 ints
	var y = [3]int{10, 20, 30}
	// [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	var a = [...]int{10, 20, 30}

	fmt.Println(len(x))
	fmt.Println(len(y))
	fmt.Println(len(z))
	fmt.Println(len(a))
}
