package main

import "fmt"

func main() {

	// array := [...]string{"Ibra", "Angel", "Baim", "Yaya"}

	// slice1 := array[0:2]
	// fmt.Println(slice1)

	// slice2 := array[:2]
	// fmt.Println(slice2)

	// slice3 := array[2:]
	// fmt.Println(slice3)

	// slice4 := array[:]
	// fmt.Println(slice4)

	// angka := []int{10, 20, 30}
	// fmt.Println(len(angka))

	days := [7]string{
		"Senin",
		"Selasa",
		"Rabu",
	}

	slice1 := days[:]
	fmt.Println(slice1)

	append1 := append(slice1, "Kamis")
	fmt.Println(append1)
	fmt.Println(days)
}
