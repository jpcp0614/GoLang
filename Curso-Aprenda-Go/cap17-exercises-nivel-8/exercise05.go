/*
Partindo do código abaixo, ordene os []user por idade e sobrenome.

- Os valores no campo Sayings devem ser ordenados também, e demonstrados de maneira esteticamente harmoniosa.
*/

package main

import (
	"fmt"
	"sort"
)

type User struct {
	First			string
	Last			string
	Age				int
	Sayings		[]string
}

type ordenarPorIdade []User
type ordenarPorSobrenome []User

func (a ordenarPorIdade) Len() int { return len(a) }
func (a ordenarPorIdade) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ordenarPorIdade) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ordenarPorSobrenome) Len() int { return len(a) }
func (a ordenarPorSobrenome) Less(i, j int) bool { return a[i].Last < a[j].Last }
func (a ordenarPorSobrenome) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func main() {
	u1 := User{
		First: "James",
		Last: "Bond",
		Age: 32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last: "Moneypenny",
		Age: 27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last: "Hmmmm",
		Age: 54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(ordenarPorIdade(users))

	for _, v := range users {
		fmt.Println("Nome:", v.First, "\t\tSobrenome:", v.Last, "\t\tIdade:", v.Age, "\nSayings:")
		for i, v2 := range v.Sayings {
			fmt.Println(i, "\t", v2)
		}
	}


	// sort.Sort(ordenarPorSobrenome(users))
	// fmt.Println(users)

}