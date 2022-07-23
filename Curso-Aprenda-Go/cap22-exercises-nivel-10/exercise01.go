/*
Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
    - Usando uma função anônima auto-executável
    - Usando buffer
*/

package main

import (
	"fmt"
)

func main() {
	// com goroutines
	// c := make(chan int)
	// go func() {
	// 	c <- 42
	// }()

	// com buffer
	c := make(chan int, 1)
	c <- 42

	fmt.Println(<-c)
}