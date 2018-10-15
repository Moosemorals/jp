// jp is a command line tool to prettyfy json
package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func format(in io.Reader, out io.Writer) {
	var v interface{}

	decoder := json.NewDecoder(in)
	err := decoder.Decode(&v)
	if err != nil {
		log.Fatalf("Can't read JSON: %+v", err)
	}

	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(&v)
	if err != nil {
		log.Fatalf("Can't write JSON: %+v", err)
	}
}

func main() {
	format(os.Stdin, os.Stdout)
}
