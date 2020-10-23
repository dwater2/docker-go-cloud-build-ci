package main

import "fmt"

func Soma(a int, b int) int {
	retorno := a + b
	fmt.Printf("%d + %d = %d\n", a, b, retorno)
    return retorno
}

func main() {
	a := 5
    b := 5
	Soma(a, b)
}