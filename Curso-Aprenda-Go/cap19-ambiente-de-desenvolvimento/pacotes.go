/*
Opção 1: uma pasta, vários arquivos.
    - package declaration em todos os arquivos
    - package scope: um elemento de um arquivo é acessível de todos os arquivos
    - imports tem file scope
- Opção 2: separando por packages.
    - pastas diferentes
    - requer imports
    - para usar: package.Função()
- Exportado vs. não-exportado, ou seja, visível vs. não-visível
    - Em Go não utilizamos os termos "público" e "privado" como em outras linguagens
    - É somente questão de capitalização
        - Com maiúscula: exportado, visível fora do package
        - Com minúscula: não exportado, não utilizável fora do package
*/

package main

import (
  "fmt"
)

func main() {
  fmt.Println()
}