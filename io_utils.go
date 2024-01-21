// Modify the text received by a unix pipe to write it to Stdout.
//
// José Almuedo [https://github.com/jalmuedo]
// This code is in the public domain.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func isInputFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode()&os.ModeCharDevice == 0
}

func digestIo(r io.Reader, w io.Writer) error {
	// By default bufio uses 4096 bytes long buffer.
	//
	// In Linux versions before 2.6.11, the capacity of a pipe was the same as
	// the system page size (e.g., 4096 bytes on i386). Since Linux 2.6.11, the
	// pipe capacity is 65536 bytes. Since Linux 2.6.35, the default pipe
	// capacity is 65536 bytes.
	//
	// For me it's ok with the default value, but you can change the buffer
	// size to increase performance.
	//
	// scanner := bufio.NewScanner(bufio.NewReaderSize(r, 64*1024))

	scanner := bufio.NewScanner(bufio.NewReader(r))

	for scanner.Scan() {
		text := colorizeChunks(scanner.Text(), os.Args[1:])
		_, err := fmt.Fprintln(w, text)
		if scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
