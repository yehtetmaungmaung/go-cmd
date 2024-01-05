package main

import (
	"bytes"
	"testing"
)

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 3
	res := count(b, true, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 6
	res := count(b, false, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 32
	res := count(b, false, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
