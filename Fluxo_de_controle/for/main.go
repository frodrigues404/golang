package main

import "fmt"

func main() {
	texto := "Texto"
	for i := 0; i < len(texto); i++{
		if string(texto[i]) == "t" { break };
		fmt.Println(string(texto[i]))
	}

	// While sample
	i := true;
	for i == true {	fmt.Println("1"); }
}