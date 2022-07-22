/*
Quando temos uma slice, podemos passar os elementos individuais através "deste..." operador.
- Exemplos:
    - Desenrolando uma slice de ints com como argumento para a função "soma" anterior
        - Go Playground: https://play.golang.org/p/k8O3__8UDa
    - Pode-se passar *zero* ou mais valores
        - Go Playground: https://play.golang.org/p/C238I9n7Vs
    - O parâmetro variádico deve ser o parâmetro final → ref/spec#Passing_arguments_to_..._parameters
*/

package main

import (
	"fmt"
)

func main() {

	sliceEvents := []int{1, 2, 3, 4, 5}

	total1 := multParam(sliceEvents...)
	total2 := multParam(1, 2, 3, 4, 5)
	total3 := multParam()

	fmt.Println(total1)
	fmt.Println(total2)
	fmt.Println(total3)

}

func multParam(x ...int) (int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}