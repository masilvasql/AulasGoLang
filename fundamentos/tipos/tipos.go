package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 3, 1000)
	fmt.Println("Literal inteiro é ", reflect.TypeOf(32000))

	//sem sinal (só positivos) uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("o byte é ", reflect.TypeOf(b))

	// com sinal int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("o valor máximo do int é ", i1)
	fmt.Println("o tipo é ", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa o valor inteiro da tabela UNICODE(int32)
	fmt.Println("o valor de a é: ", i2)
	fmt.Println("o tipo dele é", reflect.TypeOf(i2))

	// numeros reais (float32) e (float64)
	var x float32 = 49.99
	fmt.Println("O TIPO é ", reflect.TypeOf(x))
	fmt.Println("o tipo literal de 49.99 é ", reflect.TypeOf(49.99))

	bo := true
	fmt.Println("o tipo é", reflect.TypeOf(bo))

	//string
	s1 := "Olá, sou uma string"
	fmt.Println("meu tipo é ", reflect.TypeOf(s1))
	fmt.Println("Meu tamanho é", len(s1))

	//string com multiplas linhas

	s2 := `Olá 
	meu 
	nome
	é
	Marcelo`

	fmt.Println("Olá ", s2)

	//char ?
	char := 'a'
	fmt.Println("tipo de char é", reflect.TypeOf(char))
}
