package main

import "fmt"

func main() {
	// versi klasik
	var Nama string = "Rashad"
	var Umur int = 17

	// versi golang menebak type data
	Kota := "Jakarta"
	tahun := 2009

	// Contanta
	const pi = 3.14 
	const AppName = "App Manajer"
	// Constanta ini tidak bisa dirubah nilainya, kalo diubah bakal error
	
	fmt.Println("Pi:",pi)
	fmt.Println("Nama Aplikasi:", AppName)

	fmt.Println(Nama)
	fmt.Println(Umur)
	fmt.Println(Kota)
	fmt.Println(tahun)

	// Contoh kalo diubahh
	//pi = 3.1415 -> ini bakal error
}