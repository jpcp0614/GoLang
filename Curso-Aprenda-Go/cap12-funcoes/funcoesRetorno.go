/*
Pode-se usar uma função como retorno de uma função
- Declaração: func f() return
- Exemplo: func f() func() int { [...]; return func() int{ return [int] } }
    - ????: fmt.Println(f()())
*/

package main

import(
	"fmt"
)

// func main() {
// 	f := retornaFunc(10)
// 	fmt.Println(f(20))
// }

// func retornaFunc(x int) func(int) int {
// 	return func(y int) int {
// 		return x + y
// 	}
// }

func main() {
	f := retornaFunc()
	fmt.Println(f(20))

	// outra forma
	fmt.Println(retornaFunc()(20))
}

func retornaFunc() func(int) int {
	return func(y int) int {
		return 40 + y
	}
}