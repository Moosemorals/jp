// jp is a command line tool to prettyfy json
package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

func format(opts options, in io.Reader, out io.Writer) {
	var v interface{}

	decoder := json.NewDecoder(in)
	err := decoder.Decode(&v)
	if err != nil {
		log.Fatalf("Can't read JSON: %+v", err)
	}

	encoder := json.NewEncoder(out)
	if !opts.compact {
		encoder.SetIndent("", strings.Repeat(" ", opts.indent))
	}
	err = encoder.Encode(&v)
	if err != nil {
		log.Fatalf("Can't write JSON: %+v", err)
	}
}

type options struct {
	compact bool
	indent  int
}

func main() {
	var opts options

	flag.BoolVar(&opts.compact, "compact", false, "Set to remove indents")
	flag.IntVar(&opts.indent, "indent", 2, "Width of indent")
	help := flag.Bool("help", false, "Show this help")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	format(opts, os.Stdin, os.Stdout)
}
