package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"
)

func TestAssembler(t *testing.T) {
	same := deepCompare("example.hack", "result.hack")
	if !same {
		t.Fail()
	}
}

func deepCompare(file1, file2 string) bool {
	f1, err1 := ioutil.ReadFile(file1)

	if err1 != nil {
		log.Fatal(err1)
	}

	f2, err2 := ioutil.ReadFile(file2)

	if err2 != nil {
		log.Fatal(err2)
	}

	return bytes.Equal(f1, f2)
}
