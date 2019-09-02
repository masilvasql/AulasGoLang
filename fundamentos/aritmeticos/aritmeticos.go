package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("subração => ", a-b)
	fmt.Println("divisão => ", a/b)
	fmt.Println("multiplicação => ", a*b)
	fmt.Println("módulo => ", a%b)

	//bitwise

	fmt.Println("E =>", a&b)    //valor binario 11 & 10 = 10
	fmt.Println("OU => ", a|b)  //valor binario 11 | 10 = 11
	fmt.Println("Xor => ", a^b) //valor binario 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	//outras operações usando math
	fmt.Println("maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("exponencial =>", math.Pow(c, d))

}
