/*
Range: for k, v := range map { }
- Reiterando: maps *não tem ordem* e um range usará uma ordem aleatória.
- Go Playground: https://play.golang.org/p/6zEMfIP-AE
- delete(map, key)
- Deletar uma key não-existente não retorna erros!
*/

package main

import (
	"fmt"
)

func main() {
	qualquerCoisa := map[int]string{
		123: "É muito legal",
		982: "Menos legal",
		215: "Esse é massa",
		18: "Serve pra alguma coisa",
	}

	// map[18:Serve pra alguma coisa 123:É muito legal 215:Esse é massa 982:Menos legal]
	fmt.Println(qualquerCoisa)

	for key, value := range qualquerCoisa {
		fmt.Println(key, value) // não sai na mesma ordem
	}

	totalKey := 0

	for key := range qualquerCoisa {
		totalKey += key
	}
	fmt.Println(totalKey) // 1338

	// deletar no mapa
	delete(qualquerCoisa, 18)

	// map[123:É muito legal 215:Esse é massa 982:Menos legal]
	fmt.Println(qualquerCoisa)
}