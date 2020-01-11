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
	generateHeader(writer)
	generateBody(writer)
	generateReference(writer)
	generateTrailer(writer)
}

func generateHeader(writer io.Writer) {
	fmt.Fprintln(writer, "%PDF-1.7")
	fmt.Fprintln(writer, "%????")
}

func generateBody(writer io.Writer) {
	fmt.Fprintln(writer, "1 0 obj")
	fmt.Fprintln(writer, "<<")
	fmt.Fprintln(writer, "/Pages 2 0 R")
	fmt.Fprintln(writer, "/Type /Catalog")
	fmt.Fprintln(writer, ">>")
	fmt.Fprintln(writer, "endobj")
	fmt.Fprintln(writer, "2 0 obj")
	fmt.Fprintln(writer, "<<")
	fmt.Fprintln(writer, "/Kids [3 0 R]")
	fmt.Fprintln(writer, "/Count 1")
	fmt.Fprintln(writer, "/Type /Pages")
	fmt.Fprintln(writer, ">>")
	fmt.Fprintln(writer, "endobj")
	fmt.Fprintln(writer, "3 0 obj")
	fmt.Fprintln(writer, "<<")
	fmt.Fprintln(writer, "/Parent 2 0 R")
	fmt.Fprintln(writer, "/MediaBox [0 0 595 842]")
	fmt.Fprintln(writer, "/Resources 4 0 R")
	fmt.Fprintln(writer, "/Contents 5 0 R")
	fmt.Fprintln(writer, "/Type /Page")
	fmt.Fprintln(writer, ">>")
	fmt.Fprintln(writer, "endobj")
	fmt.Fprintln(writer, "4 0 obj")
	fmt.Fprintln(writer, "<<")
	fmt.Fprintln(writer, "/Font")
	fmt.Fprintln(writer, "<<")
	fmt.Fprintln(writer, "/F0")
	fmt.Fprintln(writer, "<<")
	fmt.Fprintln(writer, "/BaseFont /Times-Roman")
	fmt.Fprintln(writer, "/Subtype /Type1")
	fmt.Fprintln(writer, "/Type /Font")
	fmt.Fprintln(writer, ">>")
	fmt.Fprintln(writer, ">>")
	fmt.Fprintln(writer, ">>")
	fmt.Fprintln(writer, "endobj")
	fmt.Fprintln(writer, "5 0 obj")
	fmt.Fprintln(writer, "<<")
	fmt.Fprintln(writer, "/Length 59")
	fmt.Fprintln(writer, ">>")
	fmt.Fprintln(writer, "stream")
	fmt.Fprintln(writer, "1. 0. 0. 1. 50. 750. cm")
	fmt.Fprintln(writer, "BT")
	fmt.Fprintln(writer, "/F0 20 Tf")
	fmt.Fprintln(writer, "(Hello, world!) Tj T*")
	fmt.Fprintln(writer, "ET")
	fmt.Fprintln(writer, "endstream")
}

func generateReference(writer io.Writer) {
	fmt.Fprintln(writer, "xref")
	fmt.Fprintln(writer, "0 6")
	fmt.Fprintln(writer, "0000000000 65535 f")
	fmt.Fprintln(writer, "0000000015 00000 n")
	fmt.Fprintln(writer, "0000000066 00000 n")
	fmt.Fprintln(writer, "0000000223 00000 n")
	fmt.Fprintln(writer, "0000000125 00000 n")
	fmt.Fprintln(writer, "0000000329 00000 n")
}

func generateTrailer(writer io.Writer) {
	fmt.Fprintln(writer, "trailer")
	fmt.Fprintln(writer, "<<")
	fmt.Fprintln(writer, "/Root 1 0 R")
	fmt.Fprintln(writer, "/Size 6")
	fmt.Fprintln(writer, ">>")
	fmt.Fprintln(writer, "startxref")
	fmt.Fprintln(writer, "0")
	fmt.Fprintln(writer, "%%"+"EOF")
}
