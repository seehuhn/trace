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
	"testing"
	"time"
)

type test struct {
	path string
	prio Priority
	tmpl string
	args []interface{}

	shouldCall  bool
	expectedMsg string
}

func TestT(t *testing.T) {
	testData := []test{
		test{
			path:       "trace",
			prio:       PrioError,
			tmpl:       "hello",
			shouldCall: true,
		},
		test{
			path:       "trace",
			prio:       PrioInfo,
			tmpl:       "hello",
			shouldCall: true,
		},
		test{
			path:       "trace",
			prio:       PrioDebug,
			tmpl:       "hello",
			shouldCall: false,
		},

		test{
			path:       "tes",
			prio:       PrioError,
			tmpl:       "hello",
			shouldCall: false,
		},
		test{
			path:       "tracea",
			prio:       PrioError,
			tmpl:       "hello",
			shouldCall: false,
		},
		test{
			path:       "trace/a",
			prio:       PrioError,
			tmpl:       "hello",
			shouldCall: true,
		},

		test{
			path:        "trace",
			prio:        PrioError,
			tmpl:        "hello %d %d %d",
			args:        []interface{}{1, 2, 3},
			shouldCall:  true,
			expectedMsg: "hello 1 2 3",
		},
	}

	var called bool
	var seenMsg string
	handle := Register(
		func(t time.Time, path string, prio Priority, msg string) {
			called = true
			seenMsg = msg
		}, "trace", PrioInfo)

	tryOne := func(idx int, run test) {
		called = false
		if run.expectedMsg == "" {
			run.expectedMsg = run.tmpl
		}
		T(run.path, run.prio, run.tmpl, run.args...)
		if called != run.shouldCall {
			t.Errorf("%d: should have called listener: %v, did call: %v",
				idx, run.shouldCall, called)
		} else if called && seenMsg != run.expectedMsg {
			t.Errorf("expected message %q, got %q", run.expectedMsg, seenMsg)
		}
	}
	for k, run := range testData {
		tryOne(k, run)
	}
	handle.Unregister()

	tryOne(-1, test{
		path:       "trace",
		prio:       PrioError,
		tmpl:       "hello",
		shouldCall: false,
	})
}

func TestEmptyPath(t *testing.T) {
	seen := false
	handler := func(t time.Time, path string, prio Priority, msg string) {
		seen = true
	}
	handle := Register(handler, "", PrioAll)
	T("test", PrioInfo, "hello")
	handle.Unregister()
	if !seen {
		t.Error("failed to call listener")
	}
}

func handlerFunc(t time.Time, path string, prio Priority, msg string) {
	// do nothing
}

func BenchmarkFunctionCall(b *testing.B) {
	t := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handlerFunc(t, "test", PrioInfo, "hell")
	}
}

func BenchmarkNoListeners(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T("trace", PrioInfo, "hello")
	}
}

func BenchmarkOtherListeners(b *testing.B) {
	handle1 := Register(handlerFunc, "path1", PrioInfo)
	handle2 := Register(handlerFunc, "path2", PrioInfo)
	for i := 0; i < b.N; i++ {
		T("elsewhere", PrioInfo, "hello")
	}
	handle1.Unregister()
	handle2.Unregister()
}

func BenchmarkFirstListener(b *testing.B) {
	handle1 := Register(handlerFunc, "path1", PrioInfo)
	handle2 := Register(handlerFunc, "path2", PrioInfo)
	for i := 0; i < b.N; i++ {
		T("path1", PrioInfo, "hello")
	}
	handle1.Unregister()
	handle2.Unregister()
}

func BenchmarkSecondListener(b *testing.B) {
	handle1 := Register(handlerFunc, "path1", PrioInfo)
	handle2 := Register(handlerFunc, "path2", PrioInfo)
	for i := 0; i < b.N; i++ {
		T("path2", PrioInfo, "hello")
	}
	handle1.Unregister()
	handle2.Unregister()

}

func BenchmarkBothListeners(b *testing.B) {
	handle1 := Register(handlerFunc, "/trace", PrioInfo)
	handle2 := Register(handlerFunc, "/trace/a", PrioInfo)
	for i := 0; i < b.N; i++ {
		T("/trace/a/b", PrioInfo, "hello")
	}
	handle1.Unregister()
	handle2.Unregister()
}
