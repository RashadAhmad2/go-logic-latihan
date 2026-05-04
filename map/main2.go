package main

import "fmt"

func main() {
	buah := map[string]int{
		"apel":   10000,
		"jeruk":  8000,
		"mangga": 12000,
	}
	for key, value := range buah {
		if value > 9000 {
			fmt.Println(key+ ":", value)
		}
	}
}