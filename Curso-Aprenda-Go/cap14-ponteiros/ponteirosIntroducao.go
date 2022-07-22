/*
Os valores (todos) ficam armazenados na memória.
- Toda localização na memória possui um endereço.
- Um ponteiro se refere a esse endereço.
- Notações:
    - &variável mostra o endereço de uma variável
        - %T: variável vs. &variável
    - *variável faz de-reference, mostra o valor que consta nesse endereço
    - ????: *&var funciona!
    - *type é um tipo que contem o endereço de um valor do tipo type, nesse caso * não é um operador
- Exemplo: a := 0; b := &a; *b++
*/

package main

import (
	"fmt"
)

func main() {
	x := 100 // inteiro

	y := &x // ponteiro que armazena um valor num lugar da memória
	fmt.Print(x, y, "\n") // 100 0xc00001a128100

	*y++
	
	fmt.Print(x, y, "\n")


	fmt.Println(*y)
	fmt.Printf("%T - %T", x, y)
}