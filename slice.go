	package main

	import "fmt"

	func main() {
		names := []string{
			"Rashad", "Ahmad", "Radja", "Raihan", "Rizik", "Ryo",
		}
		slice := names[4:6]

		fmt.Println(slice[0])
		fmt.Println(slice[1])
	}