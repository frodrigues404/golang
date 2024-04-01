package main

import "fmt"

func main() {
	var salario float32
	var salarioBonus float32
	
	salario = 1000.00
	salarioBonus = salario + (salario * 10) / 100

	if salarioBonus >= 1100 {
		fmt.Println("Salário maior")
	} else if salarioBonus == 1213 {
		fmt.Println("afsdf")
	}else {
		fmt.Println("menor")
	}

	fmt.Println("Salário: ", salarioBonus)
}