package util

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"unicode/utf8"
)

const (
	kLexBufferSize     = 1024 * 32
	kLexByteBufferSize = kLexBufferSize * utf8.UTFMax
)

var ErrLexBufferOverFlow = errors.New("lex buffer overflow")
var ErrLexBufferUnderFlow = errors.New("lex buffer underflow")

type LexRune struct {
	r      rune
	size   int
	offset int
}

type LexReader struct {
	reader *bufio.Reader
	rbuff  [kLexBufferSize]LexRune
	bbuff  [kLexByteBufferSize]byte
	ri     int
	wi     int
	ai     int
	bri    int
	bwi    int
	bai    int
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
	rd.bri = rd.bai
}

// Accept moves forward the read point for next n runes
// If there is not enough runes to accept, it returns ErrLexBufferUnderFlow,
// and nothing changes.
func (rd *LexReader) Accept(n int) (err error) {
	if rd.ai+n > rd.wi {
		err = ErrLexBufferUnderFlow
	} else {
		rd.ai += n
		r := rd.rbuff[(rd.ai-1)%kLexBufferSize]
		rd.bai = r.offset + r.size
		rd.Reset()
	}
	return
}

// Accept moves forward the read point for next n bytesj
// If there is not enough bytes to accept, it returns ErrLexBufferUnderFlow,
// and nothing changes.
func (rd *LexReader) AcceptBytes(n int) (err error) {
	if rd.bai+n > rd.bwi {
		err = ErrLexBufferUnderFlow
	} else {
		begin := rd.bai
		left := rd.ai
		right := rd.wi
		var mid int
		for {
			mid = (left + right) / 2
			r := rd.rbuff[mid%kLexBufferSize]
			if r.offset+r.size-begin == n {
				break
			} else if r.offset+r.size-begin < n {
				if left == mid {
					panic("buff error")
				}
				left = mid
			} else {
				right = mid
			}
		}
		rd.ai = mid + 1
		rd.bai += n
		rd.Reset()
	}
	return
}

// ReadRune reads the next rune of input.
func (rd *LexReader) ReadRune() (r rune, size int, err error) {
	if rd.ri == rd.wi {
		rd.nomore = true
	}
	if rd.nomore {
		if rd.wi == rd.ai+kLexBufferSize {
			return 0, 0, ErrLexBufferOverFlow
		}

		r, size, err = rd.reader.ReadRune()
		if err == nil {
			rd.rbuff[rd.wi%kLexBufferSize] = LexRune{r, size, rd.bwi}
			rd.wi++
			rd.ri++
			if rd.bwi+utf8.UTFMax > kLexByteBufferSize {
				bs := make([]byte, utf8.UTFMax, utf8.UTFMax)
				utf8.EncodeRune(bs, r)
				for i := 0; i < size; i++ {
					rd.bbuff[(rd.bwi+i)%kLexByteBufferSize] = bs[i]
				}
			} else {
				utf8.EncodeRune(rd.bbuff[rd.bwi%kLexByteBufferSize:], r)
			}
			rd.bwi += size
			rd.bri += size
		}
	} else {
		lex_rune := rd.rbuff[rd.ri%kLexBufferSize]
		r = lex_rune.r
		size = lex_rune.size
		err = nil
		rd.ri++
		rd.bri += size
	}
	return
}

// ReadString reads the next n runes and return them as string
func (rd *LexReader) ReadString(n int) (s string, err error) {

	before := rd.bri
	for i := 0; i < n; i++ {
		_, _, err = rd.ReadRune()
		if err != nil {
			return
		}
	}
	after := rd.bri
	if before/kLexByteBufferSize == after/kLexByteBufferSize {
		s = string(rd.bbuff[before%kLexByteBufferSize : after%kLexByteBufferSize])
	} else {
		buff := bytes.Buffer{}
		for i := before; i < after; i++ {
			buff.WriteByte(rd.bbuff[i%kLexByteBufferSize])
		}
		s = buff.String()
	}
	return
}
