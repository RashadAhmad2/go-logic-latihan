package main

func subarraySum(nums []int, k int) int {
    var total_Saat_ini int
	var point int
	catatanBendera := make(map[int]int)
	catatanBendera[0] = 1
		for _, v := range nums {
			total_Saat_ini += v
		jumlah, adaGk := catatanBendera[total_Saat_ini-k]
		if adaGk == true {
			point += jumlah
		}
	catatanBendera[total_Saat_ini] += 1
}
	return point
}