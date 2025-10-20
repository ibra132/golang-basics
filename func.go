package main

import "fmt"

// Function void
func sayHello(name string) {
	fmt.Println("Hello", name)
}

// Function with return
func getHello(name string) string {
	return "Hello " + name
}

// Function with multiple return value
func getFullName() (string, string) {
	return "Ibra", "Baim"
}

// Function with named return value
func getFullName2() (firstName, middleName, lastName string) {
	firstName = "Andry"
	middleName = "Putra"
	lastName = "Hakim"

	return firstName, middleName, lastName
}

// Variadic function (parameter flexible)
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Function as value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

type Filter func(string) string

// Function as param
func sayHelloWithFilter(name string, filter Filter) string {
	return "Hello " + filter(name)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// Recursive function
func factorialLoop(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoop(value-1)
	}
}

func main() {
	// 1️⃣ Function void
	fmt.Println("=== sayHello ===")
	sayHello("Ibra")
	sayHello("Budi")

	// 2️⃣ Function with return value
	fmt.Println("\n=== getHello ===")
	message := getHello("Andi")
	fmt.Println(message)

	// 3️⃣ Function with multiple return
	fmt.Println("\n=== getFullName ===")
	firstName, lastName := getFullName()
	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)

	// 4️⃣ Function with named return value
	fmt.Println("\n=== getFullName2 ===")
	f, m, l := getFullName2()
	fmt.Println("First:", f)
	fmt.Println("Middle:", m)
	fmt.Println("Last:", l)

	// 5️⃣ Variadic function
	fmt.Println("\n=== sumAll ===")
	total1 := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Sum (1,2,3,4,5):", total1)

	numbers := []int{10, 10, 10}
	total2 := sumAll(numbers...) // menggunakan slice
	fmt.Println("Sum slice [10,10,10]:", total2)

	// 6️⃣ Function as value
	fmt.Println("\n=== Function as Value ===")
	goodBye := getGoodBye
	fmt.Println(goodBye("Ibra"))
	fmt.Println(goodBye("Tono"))

	// 7️⃣ Function as param
	filter := spamFilter
	fmt.Println("\n=== Function as Param ===")
	fmt.Println(sayHelloWithFilter("Ibra", spamFilter))
	fmt.Println(sayHelloWithFilter("Anjing", filter))

	// 8️⃣ Anonymous function
	fmt.Println("\n=== Anonymous function ===")
	blacklist := func(name string) bool {
		return name == "bahlil"
	}
	registerUser("Rio", func(name string) bool {
		return name == "bahlil"
	})
	registerUser("bahlil", blacklist)

	// 9️⃣ Recursive function
	fmt.Println("\n=== Recursive function ===")
	fmt.Println(factorialLoop(2))
}
