package util

import (
	"bytes"
	"io"
	"testing"
)

func TestLex(t *testing.T) {
	data := []byte("t测试aa bbccc \n")
	lex := NewLex(bytes.NewReader(data))
	if s, _ := lex.Peek(1); s != "t" {
		t.Error(s)
	}
	if s, _ := lex.Peek(1); s != "测" {
		t.Error(s)
	}
	if s, _ := lex.Peek(1); s != "试" {
		t.Error(s)
	}
	if s, _ := lex.Read(); s != "t测试" {
		t.Error(s)
	}
	if s, _ := lex.Peek(3); s != "aa " {
		t.Error(s)
	}
	delim := []byte("b\n")
	if s, _ := lex.ReadString(delim[0]); s != "aa b" {
		t.Error(s)
	}
	if s, _ := lex.ReadString(delim[1]); s != "bccc \n" {
		t.Error(s)
	}
	if _, err := lex.Peek(1); err != io.EOF {
		t.Error(err)
	}
}
