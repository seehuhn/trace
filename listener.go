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
	"sync"
	"time"
)

// Listener is the type of functions which can be registered using the
// Register() function.
type Listener func(t time.Time, path string, prio Priority, msg string)

// ListenerHandle is the type returned by Register().  The returned
// values can be used in Unregister() to remove previously installed
// handlers.
type ListenerHandle uint

type listenerInfo struct {
	path     string
	prio     Priority
	listener Listener
}

var (
	listenerMutex sync.RWMutex   // protects listeners and listenerIdx
	listeners                    = map[ListenerHandle]*listenerInfo{}
	listenerIdx   ListenerHandle = 1
)

// Register adds the function 'listener' to the list of functions
// receiving trace messages.
//
// The argument 'path' restricts the messages received to messages
// corresponding to the given path and its sub-paths (see the
// description of 'T' for details).  The value must neither start nor
// end in a slash.  The empty string can be used to receive trace
// messages from all paths.
//
// The listener will receive all messages of priority 'prio' and
// higher.  The value PrioAll can be used for 'prio' to receive all
// messages.  The value PrioInfo can be used to receive all messages
// for the given path which do not require familiarity with the
// program source code.
func Register(listener Listener, path string, prio Priority) ListenerHandle {
	listenerMutex.Lock()
	handle := listenerIdx
	listenerIdx += 1
	listeners[handle] = &listenerInfo{
		prio:     prio,
		path:     path,
		listener: listener,
	}
	listenerMutex.Unlock()
	return handle
}

// Unregister removes a previously installed listener.  The argument
// 'handle' must be the value returned by the corresponding call to
// Register()
func (handle ListenerHandle) Unregister() {
	listenerMutex.Lock()
	delete(listeners, handle)
	listenerMutex.Unlock()
}
