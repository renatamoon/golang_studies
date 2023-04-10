package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var(
		variavel3 string = "mesa"
		variavel4 string = "cadeira"
	)
	fmt.Println(variavel3, variavel4)

	var5, var6 := "estados unidos", "brasil"
	fmt.Println(var5, var6)

	const constant1 string = "Constante 1"
	fmt.Println(constant1)

	// invertendo valor de variaveis
	var5, var6 = var6, var5
	fmt.Println(var5, var6)
}