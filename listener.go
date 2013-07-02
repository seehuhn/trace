package trace

import (
	"time"
)

// Listener is the type of functions which can be registered using the
// Register() function.
type Listener func(t time.Time, path string, prio Priority, msg string)

// ListenerHandle is the type returned by Register().  The returned
// values can be used in Unregister() to remove previously installed
// handlers.
type ListenerHandle int

type listenerInfo struct {
	path     string
	prio     Priority
	listener Listener
}

var listeners map[ListenerHandle]*listenerInfo
var listenerIdx ListenerHandle

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
	handle := listenerIdx
	listenerIdx += 1
	listeners[handle] = &listenerInfo{
		prio:     prio,
		path:     path,
		listener: listener,
	}
	return handle
}

// Unregister removes a previously installed listener.  The argument
// 'handle' must be the value returned by the corresponding call to
// Register()
func (handle ListenerHandle) Unregister() {
	delete(listeners, handle)
}

func init() {
	listeners = map[ListenerHandle]*listenerInfo{}
}
