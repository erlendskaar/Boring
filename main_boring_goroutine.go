package main

import (
	"fmt"

	"github.com/erlendskaar/Boring/boring"
)

func main() {
	c := make(chan string)
	go boring.Boring10("Favoritt faget", c)
	for i := 0; ; i++ {
		fmt.Printf("Jeg sa: %q\n", <-c)
	}
	fmt.Println("Okei den er grei.")

}
