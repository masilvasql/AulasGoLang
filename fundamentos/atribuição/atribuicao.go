package main

import "fmt"

func main() {
	var b byte = 3

	fmt.Println(b)

	i := 3 // inferência de tipo
	i += 3 // atribuição aditiva
	i -= 3 // atribuição subtrativa
	i /= 2 // divisão
	i *= 2 //multiplicativa
	i %= 2
	fmt.Println(i)

	// primeira atribuição usa-se :=
	// segunda atribuição usa-se =
	x, y := 1, 2

	fmt.Println(x, y)

	x, y = y, x

	fmt.Println(x, y)

}
