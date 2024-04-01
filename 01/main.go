package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin);

    var numero1 int;
    scanner.Scan();
    var numero2 int;
    scanner.Scan();
    fmt.Println(sumNum(numero1, numero2));
}

func sumNum(x int, y int) int{
    return x + y;
}