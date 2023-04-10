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


## <a id="links_apps"> 💻 LINKS ÚTEIS </a>

- https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/