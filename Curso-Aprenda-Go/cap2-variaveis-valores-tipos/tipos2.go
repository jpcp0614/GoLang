/*
Estrutura básica: 
    - package main.
    - func main: é aqui que tudo começa, é aqui que tudo acaba.
    - import.
- Packages:
    - Pacotes são coleções de funções pré-prontas (ou não) que você pode utilizar.
    - Notação: pacote.Identificador. Exemplo: fmt.Println()
    - Documentação: fmt.Println.
- Variáveis: "uma variável é um objeto (uma posição na memória) capaz de reter e representar um valor ou expressão."
- Variáveis não utilizadas? Não pode: _ nelas.
- ...funções variádicas.
- Lição principal: package main, func main, pacote.Identificador.
*/

package main

import (
	"fmt"
)

// func main() {
// 	x := 16
// 	y := "strings"
// 	z := true

// 	_, erros := fmt.Println("Hello World") // print line (demonstra na tela)
// 	fmt.Println(erros)
// }

func main() {
	x := 16
	y := "string"
	z := true

	fmt.Println(x, y, z)
}