/*
Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
*/

package main

import (
	"fmt"
)

func main() {
	nomes := map[string][]string{
		"joão_pereira": {"ler", "assistir séries"},
		"valdênio_martinho": {"caminhar", "ouvir música"},
		"ana_beatriz": {"Nadar", "Dormir"},
	}

	delete(nomes, "ana_beatriz")

	for key, value := range nomes {
		fmt.Println(key)
		for i, hobby := range value {
			fmt.Println("\t", i, " - ", hobby)
		}
	}
}