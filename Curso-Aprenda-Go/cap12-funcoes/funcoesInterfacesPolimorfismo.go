/*
Em Go, valores podem ter mais que um tipo.
- Uma interface permite que um valor tenha mais que um tipo.
- Declaração: keyword identifier type → type x interface
- Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
- Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
- Esse tipo será o seu tipo *e também* o tipo da interface.

- Exemplos:
    - Os tipos *profissão1* e *profissão2* contem o tipo *pessoa*
    - Cada um tem seu método *oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()*
    - Implementam a interface *gente*
    - Ambos podem acessar o função *serhumano()* que chama o método *oibomdia()* de cada *gente*
    - Também podemos no método *serhumano()* tomar ações diferentes dependendo do tipo:
        switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }* 
    - Go Playground pré-pronto: https://play.golang.org/p/VLbo_1uE-U
    https://play.golang.org/p/zGKr7cvTPF
    - Go Playground ao vivo: 
    https://play.golang.org/p/njiKbTT20Cr
- Onde se utiliza?
    - Área de formas geométricas (gobyexample.com)
    - Sort
    - DB
    - Writer interface: arquivos locais, http request/response
- Se isso estiver complicado, não se desespere. É foda mesmo. Com tempo e prática a fluência vem.
*/

package main

import (
	"fmt"
)

type pessoa struct {
	nome				string
	sobrenome		string
	idade				int
}

type dentista struct {
	pessoa
	especialista	bool
	salario				float64
}

type arquiteto struct {
	pessoa
	construcao	string
	tamanho			float64
}

type gente interface {
	oibomdia()
}

func (d dentista) oibomdia() {
	fmt.Println("Meu nome é", d.nome, "ouve só: Bom dia!")
}

func (a arquiteto) oibomdia() {
	fmt.Println("Meu nome é", a.nome, "ouve só: Bom dia!")
}

func humano(c gente) {
	switch c.(type) {
	case dentista:
		fmt.Println("Meu nome é", c.(dentista).nome, "sou especialista:", c.(dentista).especialista, "ouve só: Bom dia!")
	case arquiteto:
		fmt.Println("Meu nome é", c.(arquiteto).nome, "contruí:", c.(arquiteto).construcao, "ouve só: Bom dia!")
	}
	
}

func main() {
	mauricio := dentista {
		pessoa: pessoa {
			nome: "mauricio",
			sobrenome: "santos",
			idade: 30,
		},
		especialista: true,
		salario: 4000,
	}

	pedro := arquiteto {
		pessoa: pessoa {
			nome: "pedro",
			sobrenome: "santos",
			idade: 30,
		},
		construcao: "casa",
		tamanho: 100.5,
	}

	humano(mauricio)
	humano(pedro)

	mauricio.oibomdia()
	pedro.oibomdia()
}