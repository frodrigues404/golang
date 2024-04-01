package main

import "fmt"

func main() {
	allList := []int{1, 2, 3, 10, 2, 15, 123, 12, 31, 40}
	lessThenFive := make([]int, 0)
	for i := 0; i < len(allList); i++ {
		if allList[i] < 5 {
			lessThenFive = append(lessThenFive, allList[i])
		}
	}
	twoNumbers := lessThenFive[:2]
	fmt.Println(twoNumbers)
	fmt.Println(lessThenFive)
}
