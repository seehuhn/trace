package trace

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var hasPaniced bool
var panicMethod string

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %#v\n", r)
		} else {
			t.Error("Callers() did not panic when it should have done")
		}
	}()
	Callers()
}

func TestCallers(t *testing.T) {
	seen := false
	helper1 := func() {
		seen = true
		lines := Callers()
		fmt.Println("stack trace:")
		for _, l := range lines {
			fmt.Println("  " + l)
		}
		if len(lines) == 0 ||
			!strings.Contains(lines[0], "trace/callers_test.go") {
			t.Error("wrong stack trace")
		}
	}
	helper2 := func() {
		helper1()
	}
	handler := func(t time.Time, path string, prio Priority, msg string) {
		helper2()
	}

	handle := Register(handler, "", PrioAll)
	T("test", PrioInfo, "hello you!")
	handle.Unregister()

	if !seen {
		t.Error("failed to call listener")
	}
}
