package util

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
)

const kLexBufferSize = 1024 * 32

var ErrLexBufferOverFlow = errors.New("lex buffer overflow")
var ErrLexBufferUnderFlow = errors.New("lex buffer underflow")

type LexRune struct {
	r    rune
	size int
}

type LexReader struct {
	reader *bufio.Reader
	buff   [kLexBufferSize]LexRune
	ri     int
	wi     int
	ai     int
	nomore bool
}

// NewLexReader returns a new LexReader
func NewLexReader(reader io.Reader) *LexReader {
	rd := LexReader{}
	rd.reader = bufio.NewReader(reader)
	return &rd
}

// Reset moves backward the read point to the last accepted place.
func (rd *LexReader) Reset() {
	rd.nomore = false
	rd.ri = rd.ai
}

// Accept moves forward the read point for next n runes
// If there is not enough runes to accept, it returns ErrLexBufferUnderFlow,
// and nothing changes.
func (rd *LexReader) Accept(n int) (err error) {
	err = nil
	if rd.ai+n > rd.wi {
		err = ErrLexBufferUnderFlow
	} else {
		rd.ai += n
		rd.Reset()
	}
	return
}

// ReadRune reads the next rune of input.
func (rd *LexReader) ReadRune() (r rune, size int, err error) {
	if rd.ri == rd.wi {
		rd.nomore = true
	}
	fmt.Println(rd.ri, " ", rd.wi, " ", rd.ai)
	if rd.nomore {
		if rd.wi > rd.ri+kLexBufferSize {
			return 0, 0, ErrLexBufferOverFlow
		}

		r, size, err = rd.reader.ReadRune()
		if err == nil {
			rd.buff[rd.wi%kLexBufferSize] = LexRune{r, size}
			rd.wi++
		}
	} else {
		lex_rune := rd.buff[rd.ri%kLexBufferSize]
		r = lex_rune.r
		size = lex_rune.size
		err = nil
		rd.ri++
	}
	return
}

// ReadString reads the next n runes and return them as string
func (rd *LexReader) ReadString(n int) (s string, err error) {
	buff := bytes.Buffer{}

	defer func() {
		s = buff.String()
	}()

	bs := make([]byte, utf8.UTFMax, utf8.UTFMax)
	for i := 0; i < n; i++ {
		var r rune
		var size int
		r, size, err = rd.ReadRune()
		if err != nil {
			return
		}
		utf8.EncodeRune(bs, r)
		buff.Write(bs[:size])
	}
	return
}
