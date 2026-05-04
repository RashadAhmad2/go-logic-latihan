package main

import "fmt"
func main(){
	names := []string{
		"Rashad","Ahmad","Radja","Raihan","Rizik","Ryo",
	}
	name := names[1:5]
		fmt.Println(name)
		// loop
		for i := 0; i < len(name); i++ {
			fmt.Println(name[i])
		}
}