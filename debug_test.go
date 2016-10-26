package util

import (
	"testing"
)

func ExampleDebug() {
	SetDebugLevel(DebugOff)
	Debug("Test Debug/Null => KO")

	SetDebugLevel(DebugFatal)
	Debug("Test Debug/Fatal => KO")

	SetDebugLevel(DebugError)
	Debug("Test Debug/Error => KO")

	SetDebugLevel(DebugWarning)
	Debug("Test Debug/Warning => KO")

	SetDebugLevel(DebugInfo)
	Debug("Test Debug/Info => KO")

	SetDebugLevel(DebugDebug)
	Debug("Test Debug/Debug => OK")
	// Output:
	// [ Debug   ]: Test Debug/Debug => OK
}

func ExampleInfo() {
	SetDebugLevel(DebugOff)
	Info("Test Info/Null => KO")

	SetDebugLevel(DebugFatal)
	Info("Test Info/Fatal => KO")

	SetDebugLevel(DebugError)
	Info("Test Info/Error => KO")

	SetDebugLevel(DebugWarning)
	Info("Test Info/Warning => KO")

	SetDebugLevel(DebugInfo)
	Info("Test Info/Info => OK")

	SetDebugLevel(DebugDebug)
	Info("Test Info/Debug => OK")
	// Output:
	// [ Info    ]: Test Info/Info => OK
	// [ Info    ]: Test Info/Debug => OK
}

func ExampleWarning() {
	SetDebugLevel(DebugOff)
	Warning("Test Warning/Null => KO")

	SetDebugLevel(DebugFatal)
	Warning("Test Warning/Fatal => KO")

	SetDebugLevel(DebugError)
	Warning("Test Warning/Error => KO")

	SetDebugLevel(DebugWarning)
	Warning("Test Warning/Warning => OK")

	SetDebugLevel(DebugInfo)
	Warning("Test Warning/Info => OK")

	SetDebugLevel(DebugDebug)
	Warning("Test Warning/Debug => OK")
	// Output:
	// [ Warning ]: Test Warning/Warning => OK
	// [ Warning ]: Test Warning/Info => OK
	// [ Warning ]: Test Warning/Debug => OK
}

func ExampleError() {
	SetDebugLevel(DebugOff)
	Error("Test Error/Null => KO")

	SetDebugLevel(DebugFatal)
	Error("Test Error/Fatal => KO")

	SetDebugLevel(DebugError)
	Error("Test Error/Error => OK")

	SetDebugLevel(DebugWarning)
	Error("Test Error/Warning => OK")

	SetDebugLevel(DebugInfo)
	Error("Test Error/Info => OK")

	SetDebugLevel(DebugDebug)
	Error("Test Error/Debug => OK")

	// Output:
	// [ Error   ]: Test Error/Error => OK
	// [ Error   ]: Test Error/Warning => OK
	// [ Error   ]: Test Error/Info => OK
	// [ Error   ]: Test Error/Debug => OK
}

func TestLogWithoutSetDebugLevel(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("panic occurs, check if stream is nil")
		}
	}()

	// reset debugger
	if gDebug.stream != nil {
		gDebug.stream.Close()
	}
	gDebug = debugger{}
	Debug("")
	Info("")
	Warning("")
	Error("")
}
