// Redirect Stdin to Stdout by coloring some words received as argument.
//
// José Almuedo [https://github.com/jalmuedo]
// This code is in the public domain.
package main

import (
	"os"
)

func main() {
	if isInputFromPipe() {
		digestIo(os.Stdin, os.Stdout)
	}
}
