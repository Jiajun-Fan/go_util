package util

import (
	"bytes"
	"io"
	"regexp"
	"testing"
)

func TestLexReader(t *testing.T) {
	data := []byte("t测试aa bb好好好 \n")
	lex := NewLexReader(bytes.NewReader(data))
	if err := lex.Accept(1); err != ErrLexBufferUnderFlow {
		t.Error(err)
	}
	if s, _ := lex.ReadString(1); s != "t" {
		t.Error(s)
	}
	if s, _ := lex.ReadString(1); s != "测" {
		t.Error(s)
	}
	if s, _ := lex.ReadString(1); s != "试" {
		t.Error(s)
	}
	lex.Reset()
	if s, _ := lex.ReadString(3); s != "t测试" {
		t.Error(s)
	}
	if err := lex.Accept(3); err != nil {
		t.Error(err)
	}
	lex.Reset()
	if s, _ := lex.ReadString(3); s != "aa " {
		t.Error(s)
	}
	lex.Reset()
	re1 := regexp.MustCompile("(a+)")
	re2 := regexp.MustCompile(" (b+)(好+)")
	group := re1.FindReaderSubmatchIndex(lex)
	if group == nil || len(group) != 4 {
		t.Error(group)
	}
	if err := lex.AcceptBytes(group[1]); err != nil {
		t.Error(err)
	}
	group = re2.FindReaderSubmatchIndex(lex)
	if group == nil || len(group) != 6 {
		t.Error(group)
	}
	if err := lex.AcceptBytes(group[1]); err != nil {
		t.Error(err)
	}
	if s, _ := lex.ReadString(1); s != " " {
		t.Error(s)
	}
	if s, _ := lex.ReadString(1); s != "\n" {
		t.Error(s)
	}
	if _, err := lex.ReadString(1); err != io.EOF {
		t.Error(err)
	}
}
