package main

import "fmt"

func main() {
	list := []int{5, 10}
	sum := 0 
	for i := 0; i < len(list); i++ {
		sum += list[i];
	}
	fmt.Println(sum)
}