package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	println(notaFinal)

	fmt.Println("Teste " + string(97))

	//inteiro para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para inteiro
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("1")
	if b {
		fmt.Println("Verdadeiro")
	}

}
