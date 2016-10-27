package util

import (
	"bytes"
	"io"
	"testing"
)

func TestLex(t *testing.T) {
	data := []byte("t测试 成功了庆祝一下\n")
	lex := NewLex(bytes.NewReader(data))
	if s, _ := lex.Peek(); s != "t" {
		t.Error()
	}
	if s, _ := lex.Peek(); s != "测" {
		t.Error()
	}
	if s, _ := lex.Peek(); s != "试" {
		t.Error()
	}
	if lex.Buffer() != "t测试" {
		t.Error()
	}
	lex.Next()
	delim := []string{
		" ",
		"",
		"了",
		"庆祝",
		"\n",
	}
	if s, _ := lex.PeekUntil(delim); s != " " {
		t.Error()
	}
	if s, _ := lex.PeekUntil(delim); s != "成功了" {
		t.Error()
	}
	if s, _ := lex.PeekUntil(delim); s != "庆" {
		t.Error()
	}
	if s, _ := lex.PeekUntil(delim); s != "祝一下\n" {
		t.Error()
	}
	if s, err := lex.Peek(); err != io.EOF || s != "" {
		t.Error()
	}
	if lex.Buffer() != " 成功了庆祝一下\n" {
		t.Error()
	}
	lex.Next()
}
