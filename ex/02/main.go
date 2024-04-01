package main

import "fmt"

func main() {
	slice := [10]int{1, 2, 3, 10, 2, 15, 123, 12, 31, 40}
	lessThenFive := make([]int, 0)
	moreThenFive := make([]int, 0)
	for i := 0; i < len(slice); i++ {
		if slice[i] < 5 {
			lessThenFive = append(lessThenFive, slice[i])
		} else {
			moreThenFive = append(moreThenFive, slice[i])
		}
	}
	fmt.Println(lessThenFive);
	fmt.Println(moreThenFive);
}
