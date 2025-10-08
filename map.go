package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Ibra",
	}

	person["age"] = "19"

	fmt.Println(person["name"])
	fmt.Println(person["age"])

}
