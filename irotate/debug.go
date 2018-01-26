package main

import (
    "fmt"
    "os"
)

type Debugger int 

var gDebug Debugger = DebugFatal

func SetDebug(d Debugger) {
    if d < DebugNull || d > DebugDebug {
        panic("illegal debug value")
    }   
    gDebug = d 
}

const (
    DebugNull    Debugger = 0 
    DebugFatal   Debugger = 1 
    DebugError   Debugger = 2 
    DebugWarning Debugger = 3 
    DebugInfo    Debugger = 4 
    DebugDebug   Debugger = 5 
)

func (d Debugger) V(v Debugger) Debugger {
    if d >= v { 
        if v == DebugDebug {
            fmt.Fprintf(os.Stderr, "[  Debug  ]: ")
        } else if v == DebugInfo {
            fmt.Fprintf(os.Stderr, "[  Info   ]: ")
        } else if v == DebugWarning {
            fmt.Fprintf(os.Stderr, "[ Warning ]: ")
        } else if v == DebugError {
            fmt.Fprintf(os.Stderr, "[  Error  ]: ")
        } else if v == DebugFatal {
            fmt.Fprintf(os.Stderr, "[  Fatal ]: ")
        }   
        return d
    } else {
        return DebugNull
    }   
}

func (d Debugger) Printf(fmts string, args ...interface{}) {
    if d > 0 {
        fmt.Fprintf(os.Stderr, fmts, args...)
    }
}

func Fatal(fmts string, args ...interface{}) {
    gDebug.V(DebugFatal).Printf(fmts, args...)
    os.Exit(-1)
}

func Error(fmts string, args ...interface{}) {
    gDebug.V(DebugError).Printf(fmts, args...)
}

func Warning(fmts string, args ...interface{}) {
    gDebug.V(DebugWarning).Printf(fmts, args...)
}

func Info(fmts string, args ...interface{}) {
    gDebug.V(DebugInfo).Printf(fmts, args...)
}

func Debug(fmts string, args ...interface{}) {
    gDebug.V(DebugDebug).Printf(fmts, args...)
}
