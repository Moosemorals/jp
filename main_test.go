// Tests for jp
package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestFormat(t *testing.T) {
	from, err := os.Open("fixtures/simple.in.json")
	if err != nil {
		t.Error(err)
	}
	expected, err := ioutil.ReadFile("fixtures/simple.out.json")
	if err != nil {
		t.Error(err)
	}

	var to bytes.Buffer
	format(from, &to)

	actual := to.Bytes()
	if !bytes.Equal(actual, expected) {
		t.Fatalf("\nGot\n%s\nExpected\n%s\n", actual, expected)
	}
}
