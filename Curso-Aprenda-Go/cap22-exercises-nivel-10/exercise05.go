/*
Utilizando este código: https://play.golang.org/p/YHOMV9NYKK
- ...demonstre o comma ok idiom.
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	if !ok {
		fmt.Println("Deu zebra:", v)
	} else {
		fmt.Println("Comma ok:", v)
	}

	close(c)
	
	v, ok = <-c
	if !ok {
		fmt.Println("Deu zebra:", v)
	} else {
		fmt.Println("Comma ok:", v)
	}
}
