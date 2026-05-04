package main

import (
	"fmt"
)

func main() {
	names := []string{
		"Rashad", "Ahmad", "Rashad", "Raihan", "Rizik",
	}
	var simpan bool;
	name := names[1:]
	for i := 0; i < len(name); i++ {
		if name[i] == "Rashad" {
			simpan = true
			break;
		}
	}
		if simpan {
		fmt.Println("Data Ditemukan")
	} else {
		fmt.Println("Tidak Ditemukan")
	}

}