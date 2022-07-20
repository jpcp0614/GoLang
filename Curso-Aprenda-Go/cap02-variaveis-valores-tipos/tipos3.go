/*
- := parece uma marmota (gopher) ou o punisher.
- Uso:
    - Tipagem automática
    - Só pode repetir se houverem variáveis novas 
    - != do assignment operator (operador de atribuição)
    - Só funciona dentro de codeblocks
- Terminologia:
    - keywords (palavras-chave) são termos reservados
    - operadores, operandos
    - statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões 
    - expressão → qualquer coisa que "produz um resultado"
    - scope (abrangência)
        - package-level scope
- Lição principal:
    - := utilizado pra criar novas variáveis, dentro de code blocks
    - = para atribuir valores a variáveis já existentes
*/

package main

import (
	"fmt"
)

var y = "Olá, bom dia"

func main() {
	x := 10.0

	fmt.Println(x, y)

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20, 30

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n\n", z, z)

	a := 5 + 5 // conta

	fmt.Printf("a: %v, %T\n", a, a)

	b := 10 < 10

	fmt.Printf("b: %v, %T\n", b, b)
}