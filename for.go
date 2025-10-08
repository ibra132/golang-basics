package main

import "fmt"

func main() {
	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke", counter)
	// }

	names := [3]string{"Ibra", "Angel", "Baim"}
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}
}
