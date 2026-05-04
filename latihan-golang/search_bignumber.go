package main

import "fmt"

func main() {
	var numbers = [5]int{
		90, 100, 80, 91, 1,
	}
	tg := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > tg {
			tg = numbers[i]
		}

	}
	fmt.Println(tg)

}