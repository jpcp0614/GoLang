/*
Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
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

	nomes["ana_beatriz"] = []string{"Nadar", "Dormir"}

	for key, value := range nomes {
		fmt.Println(key)
		for i, hobby := range value {
			fmt.Println("\t", i, " - ", hobby)
		}
	}
}