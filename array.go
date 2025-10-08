package main

import "fmt"

func main() {
	var some [5]int
	some[0] = 1

	fmt.Println(some)

	var array = [5]int{
		1,
		2,
		3,
		4,
		5,
	}

	array[4] = 10

	fmt.Println(array)

	array2 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}

	fmt.Println(len(array2))
}
