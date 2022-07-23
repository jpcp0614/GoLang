package main

import (
	"errors"
	"log"
	"math"
	"fmt"
)

var ErrNorgateMath = errors.New("não é possível raiz de um número negativo")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath)

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return math.Sqrt(f), nil
}

/*
*errors.errorString
2022/07/23 13:31:44 não é possível raiz de um número negativo
exit status 1
*/