package main

func findTheDifference(s string, t string) byte {
	var hurufExtra int
	gudang  := map[rune]int{}
    for _, jumlah := range s {
		gudang[jumlah]++
			if gudang[jumlah] == -1 {
			hurufExtra = gudang[jumlah]
		}
	}
	for _, v := range t {
		gudang[v]--
		if gudang[v] == -1 {
			hurufExtra = gudang[v]
		}
	}
	return byte(hurufExtra)
}