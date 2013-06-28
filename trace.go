package trace

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type Priority int32

const (
	// PrioCritical indicates a one-line message emitted just before
	// the program has to be aborted because of an internal error.
	// The message should contain information which may help to
	// determine the underlying problem and should be phrased in a way
	// that the text makes sense to a person not familiar with the
	// source code of the program.  A message of priority PrioCritical
	// could, for example, give the name of a missing but required
	// configuration file.
	PrioCritical = Priority(2000)

	// PrioError indicates a non-fatal, one-line message which is
	// likely to be of interest to an administrator of the system
	// running the program.  The message should be phrased in a way
	// that the text makes sense to a person not familiar with the
	// source code of the program.  A message of priority PrioError
	// could, for example, indicate that the program runs with reduced
	// functionality because of a configuration error.
	PrioError = Priority(1000)

	// PrioInfo indicates one-line status messages which allow to
	// track the activity of the program, and which may be of interest
	// to a person trying to understand the normal operation of the
	// program.  The message should be phrased in a way that the text
	// makes sense to a person not familiar with the source code of
	// the program.  A message of priority PrioInfo could, for
	// example, indicate that a configuration file has been read.
	PrioInfo = Priority(0)

	// PrioDebug indicates a one-line message which is likely to be of
	// interest to a developer of the program.  The message text may
	// assume that the reader is familiar with the source code of the
	// program.  A message of priority PrioDebug could, for example,
	// indicate that a library returned an unexpected error code.
	PrioDebug = Priority(-1000)

	// PrioDebug indicates a message which may be of interest to a
	// developer of the program.  The message text may assume that the
	// reader is familiar with the source code of the program, and the
	// text may consist of several lines.  A message of priority
	// PrioDebug could, for example, give the contents of a remote
	// server response to assist with debugging of network protocol
	// incompatibility.
	PrioVerbose = Priority(-2000)

	// PrioAll is used to register a listener which receives all
	// messages for a given path.
	PrioAll = Priority(math.MinInt32)
)

// T is used to send a trace message and to the registered listeners.
//
// The argument 'path' indicates which component of the program the
// caller of T belongs to; the value consists of slash separated,
// hierarchical fields where the first field, by convention, should
// coincide with the package name.  'path' must not be the empty
// string and must neither start nor end with a slash.
//
// The argument 'prio' indicates the priority of the message, higher
// values indicating higher importance.  Messages with positive
// priority values (corresponding to the pre-defined priorities
// PrioCritical, PrioError, and PrioInfo) should be phrased in a way
// that they make sense to a person not familiar with the source code
// of the program.  Messages of priority PrioDebug or higher should
// consist of a single line.
func T(path string, prio Priority, tmpl string, args ...interface{}) {
	if len(listeners) == 0 {
		return
	}

	var t time.Time
	var msg string
	first := true
	for _, c := range listeners {
		if prio >= c.prio && strings.HasPrefix(path, c.path) {
			if l := len(c.path); l > 0 && len(path) > l && path[l] != '/' {
				continue
			}
			if first {
				t = time.Now()
				msg = fmt.Sprintf(tmpl, args...)
			}
			c.listener(t, path, prio, msg)
		}
	}
}
