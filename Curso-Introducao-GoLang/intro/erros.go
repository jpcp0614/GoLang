// Erros
// é primeira classe
// tem que tratar muito bem

package main

import "github.com/coderockr/nfe/transmitter"

func main() {
	response, err := transmitter.transmit(nfe, xml)
	if err != nil {
		panic(err) // tratamento de erro qualquer
	}
	result, err := transmitter.saveData(response, xml)
	if err != nil {
		panic(err) // tratamento de erro qualquer
	}
}


/*
erros.go:7:8: no required module provides package
github.com/coderockr/nfe/transmitter: go.mod
file not found in current directory or any
parent directory; see 'go help modules'

o pacote vai tratar o erro ou vai retornar o erro para ser tratado
*/
