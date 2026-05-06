package main

import "fmt"

func main() {
	nilai := map[string]float64{
		"Rashad": 90,
		"Ahmad":  85,
		"Raihan": 88,
		"Radja":  95,
		"Budi":   70,
	}

	total := 0.0

	// 1. Hitung total dulu
	for _, value := range nilai {
		total += value
	}

	// 2. Hitung rata-rata
	rataRata := total / float64(len(nilai))
	fmt.Println("Rata-rata:", rataRata)

	// 3. Bandingin
	for key, value := range nilai {
		if value > rataRata {
			fmt.Println("Di atas rata-rata:", key, value)
		}
	}
}