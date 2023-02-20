package main

import (
	"fmt"

	"rsc.io/quote"

	"github.com/CoderPraBhu/go-programs/greetings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	message := greetings.Hello("Prashant")
	fmt.Println(message)
}
