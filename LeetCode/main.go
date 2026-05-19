package main
func main()  {
	var numbers = [4]int{
		90,91,92,90,
	}
	seen := map[int]bool{}
	for _, num := range numbers{
		if seen[num] {
			seen[num] = true
		}
		
	}
	
}