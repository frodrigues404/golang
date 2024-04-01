package main

import "fmt"

func main() {
	lista := []int{1, 2, 4, 10, 100, 1000}
	lista = append(lista, 10)
	for i := 0; i < len(lista); i++ {
		fmt.Print(lista[i], " ")
	}
	fmt.Println()
	fmt.Println("--------")

	tools := []string{"Kubernetes", "Docker", "Jenkins"}
	tools2 := []string{"Kubernetes", "Docker", "Jenkins", "AWS"}
	tools = append(tools2, "2")

	for i := 0; i < len(tools); i++ {
		fmt.Print(tools[i], " ")
	}
	fmt.Println()
	fmt.Println("--------")
	for i := 0; i < len(tools2); i++ {
		fmt.Print(tools2[i], " ")
	}
	fmt.Println()

	listaInicial := make([]int, 10)
	fmt.Printf("%T", listaInicial)
}
