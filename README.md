# go_util

package util
    import "github.com/Jiajun-Fan/go_util"

    Package util contains some basic utilities, including

	debug
	logger
	assertion
	read config file and generate config template
	database support

FUNCTIONS

func Debug(fmts string, args ...interface{})
    Debug writes message if debug level is larger or equal than DebugDebug

func Error(fmts string, args ...interface{})
    Error writes message if debug level is larger or equal than DebugError

func Fatal(fmts string, args ...interface{})
    Fatal writes message and call osExit

func Info(fmts string, args ...interface{})
    Info writes message if debug level is larger or equal than DebugInfo

func SetDebugLevel(d DebugLevel)
    SetDebugLevel changes the debug level, default value is DebugNull it's
    safe to call this function multiple times

func Warning(fmts string, args ...interface{})
    Warning writes message if debug level is larger or equal than
    DebugWarning

TYPES

type DebugLevel int

const (
    DebugNull    DebugLevel = 0
    DebugFatal   DebugLevel = 1
    DebugError   DebugLevel = 2
    DebugWarning DebugLevel = 3
    DebugInfo    DebugLevel = 4
    DebugDebug   DebugLevel = 5
)


