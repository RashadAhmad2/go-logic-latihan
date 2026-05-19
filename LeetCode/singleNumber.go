package main

import "fmt"

func main() {
	var numbers = [5]int{
		90, 91, 92, 90, 92,
	}
	var  numss int
	seen := map[int]bool{}
	for _, v := range numbers {
		if seen[v] == true {
			seen[v] = false
		} else {
			seen[v] = true
		}
	}
	for angka, status := range seen {
		if status == true {
			numss = angka
		}
	}
	return numss
}