package util

import (
	"bufio"
	"bytes"
	"io"
)

type LexToken int

type Lex struct {
	reader      *bufio.Reader
	peeked_buff bytes.Buffer
	peeked_err  error
}

// NewLex return a new Lex.
// Caller is responsible for closing input stream.
func NewLex(r io.Reader) *Lex {
	if r == nil {
		panic("nil input stream")
	}
	reader := bufio.NewReader(r)

	return &Lex{reader, bytes.Buffer{}, nil}
}

// Peek returns the next n rune without advancing the lex.
func (lex *Lex) Peek(n int) (string, error) {

	buff := bytes.Buffer{}
	for i := 0; i < n; i++ {
		var r rune
		if r, _, lex.peeked_err = lex.reader.ReadRune(); lex.peeked_err == nil {
			if _, lex.peeked_err = buff.WriteRune(r); lex.peeked_err != nil {
				return "", lex.peeked_err
			}
		} else {
			return buff.String(), lex.peeked_err
		}
	}
	_, lex.peeked_err = lex.peeked_buff.Write(buff.Bytes())
	return buff.String(), lex.peeked_err
}

// Read returns the peeked runes and advances the lex.
func (lex *Lex) Read() (string, error) {
	defer lex.peeked_buff.Reset()
	return lex.peeked_buff.String(), lex.peeked_err
}

// ReadString reads until the first occurrence of delim in unread input.
func (lex *Lex) ReadString(delim byte) (string, error) {
	var s string
	if s, lex.peeked_err = lex.reader.ReadString(delim); lex.peeked_err != nil {
		return s, lex.peeked_err
	}

	_, lex.peeked_err = lex.peeked_buff.WriteString(s)
	return lex.Read()
}
