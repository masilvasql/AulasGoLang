package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	//fmt.Println(x / y) não funciona, precisa de conversão
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println("a nota final foi:", notaFinal)

	//cuidado...

	fmt.Println("teste " + string(123))

	//int para string
	fmt.Println("teste " + strconv.Itoa(123))

	//string para int
	//primeiro retorno é o valor, segundo é um erro se colocar um _ "underline" no segundo parâmetro, você ignora o retorno
	num, _ := strconv.Atoi("123")
	//num, err := strconv.Atoi("123")
	fmt.Println(num - 10)

	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("verdade")
	} else {
		fmt.Println("falso")
	}

}
