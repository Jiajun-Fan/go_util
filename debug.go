/*
Package util contains some basic utilities, including
	debug
	logger
	assertion
	read config file and generate config template
	database support
*/
package util

import (
	"fmt"
	"os"
)

type DebugLevel int

type debugStream interface {
	Log(d DebugLevel, msg string)
	Open()
	Close()
}

type debugger struct {
	level  DebugLevel
	stream debugStream
}

const (
	DebugNull    DebugLevel = 0
	DebugFatal   DebugLevel = 1
	DebugError   DebugLevel = 2
	DebugWarning DebugLevel = 3
	DebugInfo    DebugLevel = 4
	DebugDebug   DebugLevel = 5
)

var gDebug debugger

func initStream() {
	if gDebug.stream == nil {
		gDebug.stream = new(stdoutStream)
	}
}

// SetDebugLevel changes the debug level, default value is DebugNull
// it's safe to call this function multiple times
func SetDebugLevel(d DebugLevel) {
	if d < DebugNull || d > DebugDebug {
		panic("illegal debug value")
	}
	gDebug.level = d
}

// Fatal writes message and call osExit
func Fatal(fmts string, args ...interface{}) {
	initStream()
	gDebug.stream.Log(DebugFatal, fmt.Sprintf(fmts, args...))
	gDebug.stream.Close()
	os.Exit(-1)
}

// Error writes message if debug level is larger or equal than DebugError
func Error(fmts string, args ...interface{}) {
	initStream()
	if gDebug.level >= DebugError {
		gDebug.stream.Log(DebugError, fmt.Sprintf(fmts, args...))
	}
}

// Warning writes message if debug level is larger or equal than DebugWarning
func Warning(fmts string, args ...interface{}) {
	initStream()
	if gDebug.level >= DebugWarning {
		gDebug.stream.Log(DebugWarning, fmt.Sprintf(fmts, args...))
	}
}

// Info writes message if debug level is larger or equal than DebugInfo
func Info(fmts string, args ...interface{}) {
	initStream()
	if gDebug.level >= DebugInfo {
		gDebug.stream.Log(DebugInfo, fmt.Sprintf(fmts, args...))
	}
}

// Debug writes message if debug level is larger or equal than DebugDebug
func Debug(fmts string, args ...interface{}) {
	initStream()
	if gDebug.level >= DebugDebug {
		gDebug.stream.Log(DebugDebug, fmt.Sprintf(fmts, args...))
	}
}
