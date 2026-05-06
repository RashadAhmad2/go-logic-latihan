package main

import (
	"fmt"
	"strings"
)

func main() {
	 var  input string
	
	buah := map[string]int{
		"apel":   10000,
		"jeruk":  8000,
		"mangga": 12000,
	}

	fmt.Print("Masukkan nama buah: ")
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	value, ditemukan := buah[input]
	if ditemukan {
		fmt.Println(input,":",value)
	}else{
		fmt.Println("Buah tidak ditemukan: "+input)
	}
}

