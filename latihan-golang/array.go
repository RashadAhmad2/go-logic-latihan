package main

import "fmt"

func main() {
	var numbers = [5]int{
		10, 20, 30, 40, 50,
	}
	numbers[2] = 100
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}