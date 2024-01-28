package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Ricardo" == "Ricardo")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas", d1.Equal(d2))

	type Pessoa struct {
		Nome string
		int
	}

	p1 := Pessoa{"João", 1}
	p2 := Pessoa{"João", 2}
	fmt.Println("Pessoas", p1 == p2)
	fmt.Println(p1, p1.Nome, p1.int)
}
