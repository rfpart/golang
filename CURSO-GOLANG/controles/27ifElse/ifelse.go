package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}
func main() {
	imprimirResultado(0.3)
	imprimirResultado(5.1)
	imprimirResultado(6.2)
}
