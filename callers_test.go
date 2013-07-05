// A simple tracing framework for the Go programming language.
// Copyright (C) 2013  Jochen Voss <voss@seehuhn.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package trace

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var (
	hasPaniced  bool
	panicMethod string
)

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
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
