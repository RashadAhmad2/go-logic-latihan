package main

import "fmt"

func main() {
	numbers := []int{
		10, 20, 30, 40, 50, 60,
	}

	number := numbers[2:6]
	number[0] = 100
	fmt.Println(number)
	fmt.Println(numbers)

}