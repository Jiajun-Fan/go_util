package util

import (
	"bufio"
	"bytes"
	"io"
	"unicode/utf8"
)

type LexToken int

type Lex struct {
	scanner     *bufio.Scanner
	peeked_buff bytes.Buffer
	peeked_err  error
}

// NewLex return a new Lex.
// Caller is responsible for closing input stream.
func NewLex(r io.Reader) *Lex {
	if r == nil {
		panic("nil input stream")
	}
	s := bufio.NewScanner(bufio.NewReader(r))
	s.Split(bufio.ScanRunes)

	return &Lex{s, bytes.Buffer{}, nil}
}

// Next makes Lex move forward to next rune and return nil on success.
// If input stream rearchs the end, this function has no effect and returns IO.EOF.
func (lex *Lex) Next() error {
	if lex.peeked_err == io.EOF {
		return io.EOF
	}
	lex.peeked_buff.Reset()
	return nil
}

func (lex *Lex) Buffer() string {
	return lex.peeked_buff.String()
}

// Peek returns the next rune and error information.
// The read rune is appended to the buffer of lex.
// If input stream rearchs the end, the error information is io.EOF.
func (lex *Lex) Peek() (string, error) {
	ret := ""
	if lex.scanner.Scan() {
		ret = lex.scanner.Text()
		lex.peeked_buff.WriteString(ret)
		lex.peeked_err = nil
	} else {
		if err := lex.scanner.Err(); err == nil {
			lex.peeked_err = io.EOF
		} else {
			lex.peeked_err = err
		}
	}
	return ret, lex.peeked_err
}

// PeekUntil returns the all the runes until it reachs any string in delim,
// and error information if there is any.
// delim is an array of strings, if a string has more than 1 rune, PeekUntil only
// compare the first rune, others are ignored.
// The read runes are appended to the buffer of lex.
// If input stream rearchs the end, the error information is io.EOF.
func (lex *Lex) PeekUntil(delim []string) (string, error) {
	var buffer bytes.Buffer
	for {
		match := false
		if lex.scanner.Scan() {
			r := lex.scanner.Text()
			buffer.WriteString(lex.scanner.Text())
			lex.peeked_err = nil
			for i := range delim {
				d := delim[i]
				if utf8.RuneCountInString(d) > 0 {
					_, width := utf8.DecodeRuneInString(d)
					d = d[:width]
					if r == d {
						match = true
						break
					}
				}
			}
			if match {
				break
			}
		} else {
			if err := lex.scanner.Err(); err == nil {
				lex.peeked_err = io.EOF
			} else {
				lex.peeked_err = err
			}
			break
		}
	}
	ret := buffer.String()
	lex.peeked_buff.WriteString(ret)
	return ret, lex.peeked_err
}
