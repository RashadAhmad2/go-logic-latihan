package main

import "fmt"

func main() {
	var tg int
	var name string
	dataNilai := map[string]int{
		"Rashad": 90,
		"Ahmad":  85,
		"Raihan": 88,
		"Radja":  95,
	}
	for key, value := range dataNilai {
	if tg == 0 {
	tg = value
	name = key
	} else if value > tg {
		tg = value
		name = key
	}
	}
	fmt.Println("Nilai tertinggi adalah:", name, "dengan nilai:", tg)
}