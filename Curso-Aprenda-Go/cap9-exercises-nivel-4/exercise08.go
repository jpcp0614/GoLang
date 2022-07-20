/*
Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
*/

package main

import (
	"fmt"
)

func main() {
	nomes := map[string][]string{
		"joão_pereira": {"ler", "assistir séries"},
		"valdênio_martinho": {"caminhar", "ouvir música"},
	}

	for key, value := range nomes {
		fmt.Println(key)
		for i, hobby := range value {
			fmt.Println("\t", i, " - ", hobby)
		}
	}
}