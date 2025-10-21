package main

import "fmt"

func main() {
	counter := 0

	increment := func() int {
		counter++
		return counter
	}

	increment()
	increment()
	fmt.Println(counter)
}
