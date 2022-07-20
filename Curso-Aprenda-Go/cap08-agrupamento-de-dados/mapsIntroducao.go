/*
Utiliza o formato key:value.
- E.g. nome e telefone
- Performance excelente para lookups.
- map[key]value{ key: value }
- Acesso: m[key]
- Key sem value retorna zero. Isso pode trazer problemas.
- Para verificar: comma ok idiom.
    - v, ok := m[key]
    - ok é um boolean, true/false
- Na prática: if v, ok := m[key]; ok { }
- Para adicionar um item: m[v] = value
- Maps *não tem ordem.*
*/

package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"Rodrigo": 2222222,
		"João": 4444444,
	}

	fmt.Println(amigos) // map[João:4444444 Rodrigo:2222222]

	fmt.Println(amigos["Rodrigo"]) // 2222222

	// adiciona
	amigos["Valdênio"] = 8888888

	fmt.Println(amigos) // map[João:4444444 Rodrigo:2222222 Valdênio:8888888]

	fmt.Println(amigos["Fantasma"]) // 0

	existe, ok := amigos["Fantasma"] // 0 false

	fmt.Println(existe, ok) //

	if será, ok := amigos["Joana"]; ok {
		fmt.Println(será)
		} else {
		fmt.Println("Não tem")
	}
}