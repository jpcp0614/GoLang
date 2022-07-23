package main

import (
	"log"
	"math"
	"fmt"
)


func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("não é possível raiz de um número negativo: %v", f)
	}
	return math.Sqrt(f), nil
}

/*
2022/07/23 13:34:49 não é possível raiz de um número negativo: -10
exit status 1
*/