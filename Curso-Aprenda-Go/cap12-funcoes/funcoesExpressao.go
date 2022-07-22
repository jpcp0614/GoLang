/*
f := func(p params){ ... }
- f()
*/

package main

import (
	"fmt"
)

func main() {
	x := 10

	y := func(x int) float64 {
		return (float64(x) * .25) + 10
	}

	fmt.Println(y(x))

}