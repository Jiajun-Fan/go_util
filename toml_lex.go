package util

import ()

type Token int

const (
	TokIllegal      Token = iota // illegal
	TokSpace                     // blank \t
	TokNewLine                   // \n \r\n
	TokNumber                    // 0-9
	TokAlphabetic                // a-zA-Z
	TokUnderScore                // _
	TokDot                       // .
	TokQuota                     // "
	TokSingleQuota               // '
	TokLeftBracket               // [
	TokRightBracket              // ]
	TokHash                      // #
	TokChar                      // everything else
	TokEOF                       // end of file
)

/*type TomlLex struct {
	Lex
}

func (t *TomlLex) findSpace() {
	for r, err := t.Peek(1); err != nil; {
	}
}

func (t *TomlLex) NextToken() (string, Token) {
	lead, err := t.Peek(1)
	if err == io.EOF {
		return TokEOF
	} else {
		return TokIllegal
	}
	if r == " " || r == "\t" {
		return r, TokSpace
	}
}*/
