// Example to demonstrate [base64 stream decoder](https://golang.org/pkg/encoding/base64/#NewDecoder) on large files.
// It reads from an io.Reader and returns one that can be copied into an
// io.Writer. It also takes care of the new lines.

// This code will accept a base64 encoded file (whitespace does not matter)
// with -file/--file and store the decoded bytes in filename-out.
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// This creates a flag parameter. Allows us to call -file or --file.
	var filename string
	flag.StringVar(&filename, "file", "", "input file")
	flag.Parse()

	// It's a lie, it's pretty fast.
	fmt.Println("reading input file, this may take a while")

	// Open input file
	input, err := os.Open(filename)
	// We are panic-ing with every error because we want to stop if things
	// go wrong.
	if err != nil {
		panic(err)
	}
	// Defer close input file
	defer input.Close()

	// Open output file
	output, err := os.Create(filename + "-out")
	if err != nil {
		panic(err)
	}
	// Defer close output file
	defer output.Close()

	// Create base64 stream decoder from input file. *io.File implements the
	// io.Reader interface. In other words we can pass it to NewDecoder.
	decoder := base64.NewDecoder(base64.StdEncoding, input)
	// Output of decoder is also io.Reader, we can use io.Copy to copy it to
	// output file directly.
	io.Copy(output, decoder)

	fmt.Println("storing base64 decoder input file")
}
