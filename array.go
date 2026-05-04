package main

import "fmt"

func main() {
	var names [3]string
	// versi tidak langsung
	names[0] = "Rashad"
	names[1] = "Ahmad"
	names[2] = "Ahmad"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	// Versi Array Langsung
	var numbers = [3]int{
		90,91,92,
	}
	fmt.Print(numbers)
}