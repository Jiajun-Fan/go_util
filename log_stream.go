package util

import (
	"fmt"
)

type stdoutStream struct {
}

func (s stdoutStream) Open() {
}

func (s stdoutStream) Close() {
}

func (s stdoutStream) Log(d DebugLevel, msg string) {
	if d == DebugFatal {
		fmt.Printf("[ Fatal   ]: %s\n", msg)
	} else if d == DebugError {
		fmt.Printf("[ Error   ]: %s\n", msg)
	} else if d == DebugWarning {
		fmt.Printf("[ Warning ]: %s\n", msg)
	} else if d == DebugInfo {
		fmt.Printf("[ Info    ]: %s\n", msg)
	} else if d == DebugDebug {
		fmt.Printf("[ Debug   ]: %s\n", msg)
	}
}
