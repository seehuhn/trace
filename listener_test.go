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

func listener(t time.Time, path string, prio Priority, msg string) {
	// do nothing
}

func TestRegister(t *testing.T) {
	handle := Register(listener, "test", PrioInfo)
	if len(listeners) != 1 {
		t.Error("failed to register listener")
	}
	handle.Unregister()
	if len(listeners) != 0 {
		t.Error("failed to unregister listener")
	}
}
