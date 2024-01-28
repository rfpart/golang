package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)                                 //numeros inteiros
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000)) //pacote reflect informa o tipo da variável

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64 // vai pegar o maior valor de um inteiro de 64 bits
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal é 49.99 é", reflect.TypeOf(49.00))
	fmt.Println(x)

	// bolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Leo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linhas
	s2 := `Olá
	Meu
	nome
	é
	Leo`
	fmt.Println("O tamanho da string é", len(s2))
	fmt.Println(s2)

	// char ???   não existe char em golang
	char := 'A'
	fmt.Println("O tipo da char é", reflect.TypeOf(char))
	fmt.Println(char)
	fmt.Println(char)
}
