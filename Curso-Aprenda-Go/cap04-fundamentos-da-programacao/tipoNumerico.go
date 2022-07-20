/*
int vs. float: Números inteiros vs. números com frações.
- golang.org/ref/spec → numeric types
- Integers:
    - Números inteiros
    - int & uint → “implementation-specific sizes”
    - Todos os tipos numéricos são distintos, exceto:
        - byte = uint8
        - rune = int32 (UTF8)
        (O código fonte da linguagem Go é sempre em UTF-8).
    - Tipos são únicos
        - Go é uma linguagem estática
        - int e int32 não são a mesma coisa
        - Para "misturá-los" é necessário conversão
    - Regra geral: use somente int
- Floating point:
    - Números racionais ou reais
    - Regra geral: use somente float64
- Na prática:
    - Defaults com :=
    - Tipagem com var
    - Dá pra colocar número com vírgula em tipo int?
    - Overflow
		“implementation-specific sizes”? Runtime package. Word.
    - GOOS
    - GORUNTIME
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := "e"
	b := "é"
	c := "银"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)
	fmt.Printf("%v, %v, %v\n", d, e, f)

	fmt.Println(runtime.GOOS) // SO linux
	fmt.Println(runtime.GOARCH) // processador amd64
}

