// Redirect Stdin to Stdout by coloring some words received as argument.
//
// José Almuedo [https://github.com/jalmuedo]
// This code is in the public domain.
package main

import (
	"log"
	"os"
)

func main() {
	if isInputFromPipe() {
		err := digestIo(os.Stdin, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	}
}
