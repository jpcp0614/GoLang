/*
Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
*/

package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

type figura interface {
	area()
}

func info(f figura) {
	f.area()
}

func (l quadrado) area() {
	calc := math.Pow(l.lado, 2)
	fmt.Printf("Área do quadrado: %.2fm²\n", calc)
}

func (r circulo) area() {
	calc := 2 * r.raio * math.Pi
	fmt.Printf("Área do círculo: %.2fm²\n", calc)
}

func main() {
	quad := quadrado {
		lado: 10.,
	}
	circ := circulo {
		raio: 2.,
	}

	// quad.area()
	// circ.area()

	info(quad)
	info(circ)
}