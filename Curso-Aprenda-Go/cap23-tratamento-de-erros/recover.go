package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("Retornou de f")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("Chamando g")
	g(0)
	fmt.Println("Retornou de g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Pânico!!!")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defer no g", i)
	fmt.Println("Mostrando no g", i)
	g(i + 1)
}