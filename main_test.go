// Tests for jp
package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func testFile(t *testing.T, opts options, inFile, outFile string) {
	from, err := os.Open(inFile)
	if err != nil {
		t.Error(err)
	}
	expected, err := ioutil.ReadFile(outFile)
	if err != nil {
		t.Error(err)
	}

	var to bytes.Buffer

	format(opts, from, &to)

	actual := to.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Errorf("\nGot\n%s\nExpected\n%s\n", actual, expected)
	}
}

var testdata = []struct {
	name            string
	opts            options
	inFile, outFile string
}{
	{"false, 2", options{compact: false, indent: 2}, "fixtures/test.json", "fixtures/indented-2.json"},
	{"false, 4", options{compact: false, indent: 4}, "fixtures/test.json", "fixtures/indented-4.json"},
	{"true , 2", options{compact: true, indent: 2}, "fixtures/test.json", "fixtures/compact.json"},
}

func TestFormat(t *testing.T) {

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			testFile(t, tt.opts, tt.inFile, tt.outFile)
		})
	}

	testFile(t, options{compact: false, indent: 2}, "fixtures/test.json", "fixtures/indented-2.json")
}
