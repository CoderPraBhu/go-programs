package main

import (
	"fmt"

	"rsc.io/quote"

	"log"

	"github.com/CoderPraBhu/go-programs/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0) // time, source, line
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	message, err := greetings.Hello("Prashant")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
