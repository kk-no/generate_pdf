package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var file *os.File
	var name = "sample.pdf"

	file, _ = os.Create(name)
	defer file.Close()
	writer := io.Writer(file)
	fmt.Fprintln(writer, "%PDF-1.7")
	fmt.Fprintln(writer, "%????")
	fmt.Fprintln(writer, "startxref")
	fmt.Fprintln(writer, "0")
	fmt.Fprintln(writer, "%%"+"EOF")

}
