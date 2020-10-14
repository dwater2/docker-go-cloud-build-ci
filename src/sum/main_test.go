package main

import "testing"

func TestSoma(t *testing.T) {
	a := 5
    b := 5
    total := Soma(a, b)
    if total != 10 {
       t.Errorf("Soma incorreta, valor: %d, esperado: %d.", total, 10)
    }
}