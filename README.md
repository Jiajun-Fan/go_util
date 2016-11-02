# util
--
    import "github.com/Jiajun-Fan/go_util"

Package util contains some basic utilities, including

    debug
    logger
    assertion
    read config file and generate config template
    database support

## Usage

```go
var (
	ErrLexBufferOverFlow  = errors.New("lex buffer overflow")
	ErrLexBufferUnderFlow = errors.New("lex buffer underflow")
)
```

#### func  Debug

```go
func Debug(fmts string, args ...interface{})
```
Debug writes message if debug level is larger or equal than DebugDebug.

#### func  Error

```go
func Error(fmts string, args ...interface{})
```
Error writes message if debug level is larger or equal than DebugError.

#### func  Fatal

```go
func Fatal(fmts string, args ...interface{})
```
Fatal writes message and call os.Exit.

#### func  Info

```go
func Info(fmts string, args ...interface{})
```
Info writes message if debug level is larger or equal than DebugInfo.

#### func  SetDebugLevel

```go
func SetDebugLevel(d DebugLevel)
```
SetDebugLevel changes the debug level, default value is DebugOff. It's safe to
call this function multiple times.

#### func  Warning

```go
func Warning(fmts string, args ...interface{})
```
Warning writes message if debug level is larger or equal than DebugWarning.

#### type DebugLevel

```go
type DebugLevel int
```


```go
const (
	DebugOff DebugLevel = iota
	DebugFatal
	DebugError
	DebugWarning
	DebugInfo
	DebugDebug
)
```

#### type LexReader

```go
type LexReader struct {
}
```


#### func  NewLexReader

```go
func NewLexReader(reader io.Reader) *LexReader
```
NewLexReader returns a new LexReader.

#### func (*LexReader) Accept

```go
func (rd *LexReader) Accept(n int) (err error)
```
Accept moves forward the read point for next n runes. If there is not enough
runes to accept, it returns ErrLexBufferUnderFlow, and nothing changes.

#### func (*LexReader) AcceptBytes

```go
func (rd *LexReader) AcceptBytes(n int) (err error)
```
Accept moves forward the read point for next n bytes. If there is not enough
bytes to accept, it returns ErrLexBufferUnderFlow, and nothing changes.

#### func (*LexReader) ReadRune

```go
func (rd *LexReader) ReadRune() (r rune, size int, err error)
```
ReadRune reads the next rune of input.

#### func (*LexReader) ReadString

```go
func (rd *LexReader) ReadString(n int) (s string, err error)
```
ReadString reads the next n runes and return them as string.

#### func (*LexReader) Reset

```go
func (rd *LexReader) Reset()
```
Reset moves backward the read point to the last accepted place.

#### type Token

```go
type Token int
```


```go
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
```
