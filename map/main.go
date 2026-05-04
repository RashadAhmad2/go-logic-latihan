package main

import "fmt"

func main() {
	var input string
	dataNilai := map[string]int{
		"Rashad": 90,
		"Ahmad":  85,
		"Raihan": 88,	
	}
	// Map bisa diakses langsung dengan key tanpa perlu loop
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&input)
		value, ditemukan := dataNilai[input]
	 	if ditemukan { 
			fmt.Println(input + ":", value)

		}else{
			fmt.Println("Nama Tidak Ditemukan: "+input)
		}
}