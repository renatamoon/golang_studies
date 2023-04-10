package main

import "fmt"

// func main() {
// 	int8, int16, int32, int64
// }

// Tipos de Dados:
// Numeros -> inteiros e reais

// int, int8, int16, int32, int64 -> diferença é a quantidade de bits que eles ocupam

// func main() {
// 	var numero int8 = 100000000000000 // int8 não suporta esse numero
// 	fmt.Println(numero)
// }

// caso não especificar no int ele usa a arquitetura do seu computador para identificar qual o int

func main() {
	var numero int64 = 100000000000000 // int8 não suporta esse numero
	fmt.Println(numero)

	var numero1 := 2000
	fmt.Println(numero1)

	var numero2 uint32 = 100000
	fmt.Println(numero2)

	//alias - int32 == rune -> rune é mais usado quando mexe-se com caracteres
	var numero4 rune = 12345
	fmt.Println(numero4)

	// byte = UINT8
	var numero3 byte = 123
	fmt.Println(numero3)

	// ------------ numeros reais
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12455555555555.6
	fmt.Println(numeroReal2)

	numeroReal3 := 1234567.8
	fmt.Println(numeroReal3)

	// STRINGS
	
	var str string = "texto qualquer"
	fmt.Println(str)

	str2 := "texto qualquer 2"
	fmt.Println(str2)

	// BOLEANO

	var booleano1 bool = true
	var booleano2 bool = false

	// se quiser atribuir o valor 0 é so adicionar ela sem valor
	var booleano2 bool

	// ERRORS
	var erro error //  o valor zero dele é <nil> - 0

	var erro error = errors.New("erro interno")

}