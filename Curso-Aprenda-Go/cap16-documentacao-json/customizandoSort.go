/*
O sort que eu quero não existe. Quero fazer o meu.
- Para isso podemos usar o func Sort do package sort. Vamos precisar de um sort.Interface.
    - type Interface interface { Len() int; Less(i, j int) bool; Swap(i, j int) }
- Ou seja, se tivermos um tipo que tenha esses métodos, ao executar sort.Sort(x) as funções que vão rodar são as minhas, não as funções pré-prontas como no exercício anterior.
- E aí posso fazer do jeito que eu quiser.
- Exemplo:
    - struct carros: nome, consumo, potencia
    - slice []carros{carro1, carro2, carro3} (Sort ordena *slices!*)
    - tipo ordenarPorPotencia
    - tipo ordenarPorConsumo
    - tipo ordenarPorLucroProDonoDoPosto
*/

package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome			string
	potencia	int
	consumo		int
}

type ordenarPorNome []carro
type ordenarPorPotencia []carro
type ordenarPorConsumo []carro
type ordenarPorLucroProDonoDoPosto []carro

// ordenar por nome
func (p ordenarPorNome) Len() int { return len(p) }
func (p ordenarPorNome) Less(i, j int) bool { return p[i].nome < p[j].nome }
func (p ordenarPorNome) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Ordenar por potência
func (p ordenarPorPotencia) Len() int {	return len(p) }
func (p ordenarPorPotencia) Less(i, j int) bool {	return p[i].potencia < p[j].potencia }
func (p ordenarPorPotencia) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ordenar por consumo
func (p ordenarPorConsumo) Len() int { return len(p) }
func (p ordenarPorConsumo) Less(i, j int) bool { return p[i].consumo < p[j].consumo }
func (p ordenarPorConsumo) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ordenar por lucro para o dono do posto
func (p ordenarPorLucroProDonoDoPosto) Len() int { return len(p) }
func (p ordenarPorLucroProDonoDoPosto) Less(i, j int) bool { return p[i].consumo > p[j].consumo }
func (p ordenarPorLucroProDonoDoPosto) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	carros := []carro{
		{nome: "BMW", potencia: 320, consumo: 12},
		{nome: "Ford", potencia: 140, consumo: 14},
		{nome: "Fiat", potencia: 100, consumo: 11},
		{nome: "Honda", potencia: 200, consumo: 10},
		{nome: "Chevrolet", potencia: 180, consumo: 9},
		{nome: "VW", potencia: 220, consumo: 8},
		{nome: "Renault", potencia: 130, consumo: 7},
		{nome: "Toyota", potencia: 220, consumo: 6},
		{nome: "Mercedes", potencia: 250, consumo: 5},
	}

	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println(carros)

	sort.Sort(ordenarPorNome(carros))
	fmt.Println(carros)

	sort.Sort(ordenarPorConsumo(carros))
	fmt.Println(carros)

	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))
	fmt.Println(carros)
}





/*

DOCUMENTAÇÃO: https://pkg.go.dev/sort#Interface

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	// There are two ways to sort a slice. First, one can define
	// a set of methods for the slice type, as with ByAge, and
	// call sort.Sort. In this first example we use that technique.
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// The other way is to use sort.Slice with a custom Less
	// function, which can be provided as a closure. In this
	// case no methods are needed. (And if they exist, they
	// are ignored.) Here we re-sort in reverse order: compare
	// the closure with ByAge.Less.
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)
	*/