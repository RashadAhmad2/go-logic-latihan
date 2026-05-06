package main

import (
	"fmt"
	"math"
)

func main() {
	nilai := map[string]float64{
		"Rashad": 90,
		"Ahmad":  85,
		"Raihan": 88,
		"Radja":  95,
		"Budi":   70,
	}
	var name string
	var nilaii float64
	total := 0.0
	for _, value := range nilai {
		total += value
	}
	selisihTerkecil := 999.0
	rataRata := total / float64(len(nilai))

	for key, value := range nilai {
		selisih := math.Abs(rataRata - value)
		if selisih < selisihTerkecil {
			selisihTerkecil = selisih
			name = key
			nilaii = value
		}
		
	
	}
	fmt.Printf("Nama nilainya dekat dengan nilai rata-rata: %s\nSelisih terkecilnya: %.1f\nDengan nilai asli: %.1f\n", name, selisihTerkecil, nilaii)
}