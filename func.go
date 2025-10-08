package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Ibra", "Baim"
}

func getFullName2() (firstName, middleName, lastName string) {
	firstName = "Andry"
	middleName = "Putra"
	lastName = "Hakim"

	return firstName, middleName, lastName
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// fmt.Println(total)

	numbers := []int{10, 10, 10}
	total := sumAll(numbers...)
	fmt.Println(total)
}
