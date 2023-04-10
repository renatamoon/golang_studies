package auxiliar

import "fmt"


func Writing() {
	fmt.Println("Writing from aux package")
	writing2()  // uma função com letra minuscula somente pode ser usada dentro do mesmo pacote, ela nao pode ser chamada no arquivo main
}