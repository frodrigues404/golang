package main

import "fmt"

func main() {
	// While sample
	i := true;
	count := 1;
	for i {	
		fmt.Println(count);
		count += 1;
		if count == 10 { break; }
	}
}