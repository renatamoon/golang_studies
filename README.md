# golang_studies
Dedicated to my studies with the Golang

<p align="center">
  <a href="#aprendizado">Sobre o projeto</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#tecnologias">Tecnologias</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#comandos">Lista de Comandos</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#estrutura">Estrutura</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#links_apps">Links Úteis</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
 
</p>

## <a id="aprendizado"> 💻 APRENDIZADO </a>

Este repositório é especifico para as aulas presentes na Udemy: https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa


## <a id="tecnologias"> 💻 TECONOLOGIAS </a>

- Golang


## <a id="comandos"> 💻 LISTA DE COMANDOS </a>

- `go env` -> mostra as variáveis de ambiente do go
- `go get` + link que deseja fazer download - serve para baixar os pacotes
- `go build` -> arquivo binário executário que executa e compila tudo.
- `./modulo` -> executa o arquivo
- `go mod tidy` -> remove todos os pacotes que não estão sendo usados

## <a id="estrutura"> 💻 ESTRUTURA </a>

- bin
- pkg: pacotes baixados
- src: código fonte em go

<hr>

- Packages: grupo de arquivos que estão no mesmo diretorio e são compilados juntos. Isso quer dizer que uma variável será visivel a todos os arquivos que estão dentro do pacote. Vale para funções, constantes e arquivos.
- Imports: depois dos packages precisa ter os imports que serão usados:

```
import (	
	"fmt"
	"modulo/auxiliar"
)
```

- Função: Exemplo de função:

```
func main() {
	fmt.Println("Writing from the main file")
	auxiliar.Writing()
}
```

<hr>
- Arquivo go.mod: onde guarda os modulos junto com a versão de todos os produtos.
- tudo o que for gerado no arquivo go.mod não pode ser modificado manualmente (deve se evitar);
- se meu pacote que eu baixei for de nome: "github.com/badoux/checkmail" quando for importar dentro do meu projeto, eu so referencio o nome após a ultima barra: no caso, checkmail.
- ERRO importante no Go, caso deixar uma variavel sem usar. Ou seja, qualquer pacote ou var precisam ser usados. `variavel2 declared and not used`

<hr>
- VARIÁVEIS: podem ter duas formas de ser usadas: uma implicita ou explicita:

var variavel1 string = "Variavel 1" -> onde eu digo que a var é string
variavel2 := "Variavel 2" -> onde o proprio go compila e checa o valor da var e atribui o tipo.

- `:=` -> se chama inferencia de tipo baseado no valor dele.

- DECLARANDO MAIS DE UMA VARIAVEL:

tipando: ```var(
		variavel3 string = "mesa"
		variavel4 string = "cadeira"
	)```

inferindo: ```var5, var6 := "estados unidos", "brasil"```


## <a id="links_apps"> 💻 LINKS ÚTEIS </a>

- https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/