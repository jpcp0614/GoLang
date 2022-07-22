/*
Ponteiros permitem compartilhar endereços de memória. Isso é útil quando:
    - Não queremos passar grandes volumes de dados pra lá e pra cá
    - Queremos mudar um valor em sua localização original (tudo em Go é pass by value!)
- Exemplos:
    - x := 0; funçãoquemudaovalordoargumentopra1(x); Print(x)
    - x := 0; funçãoquemudaovalordo*argumentopra1(&x); Print(x)
*/


package main

import (
	"fmt"
)

func main() {

	x := 0

	// estaRecebeValor(x)

	estaRecebePonteiro(&x)
	fmt.Println(x)
}

func estaRecebeValor(x int) {
	x++
	fmt.Println("estaRecebeValor", x)
}

func estaRecebePonteiro(x *int) {
	*x++
	fmt.Println("estaRecebePonteiro", *x)
}