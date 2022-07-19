// Goroutines


package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}


func main() {
	go say("world") // go é quando eu uso a função (roda num segundo núcleo de processador)
	say("hello") // normal
}

/*
Ambas estão rodando ao mesmo tempo
*/