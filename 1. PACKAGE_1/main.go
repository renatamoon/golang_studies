package main


import (	
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from the main file")
	auxiliar.Writing()

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
