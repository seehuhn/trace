Trace
=====

A simple tracing framework for the Go programming language.

Copyright (C) 2013  Jochen Voss

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

The homepage of this package is at http://www.seehuhn.de/pages/trace .
Please send any comments or bug reports to the program's author,
Jochen Voss <voss@seehuhn.de> .

Overview
--------

The trace package provides a simple tracing framework for the
excellent Go programming language. Code using this framework can emit
diagnostic messages using the trace.T() function.  In normal
operation, calls to trace.T() have no effect and are very fast.  When
tracing is enabled, in order to track down a problem or to explore the
inner workings of a program, listeners can be attached to record and
display all trace messages.

Main features:
- Low overhead when no listeners are installed.
- Flexible selection of messages by listeners via message origin
  and/or priority.

Usage
-----

Code using this framework can emit diagnostic messages using the
trace.T() function.  Example:

    trace.T("a/b/c", trace.PrioError,
	    "failed to connect to server %q, using offline mode", serverName)

The first argument in this call is a path which gives information
about the origin of the message, the second argument indicates the
importance of the message.  Both, the path and the priority are
used to decide which listeners receive the correponding message.
The following arguments, a format string and additional optional
arguments, are passed to fmt.Sprintf to compose the message
reported to the listeners registered for the given message path.

Listeners can subscribe to messages, either for a given path or for
all paths, using the Register() method.  A minimum priority for
messages to be delivered can be used.  Example:

    func MyListener(t time.Time, path string, prio Priority, msg string) {
	log.Println(msg)
    }

    func main() {
	listener := trace.Register(MyListener, "a/b", trace.PrioAll)
	// ... code which calls trace.T()
	listener.Unregister()
    }

This code installs MyListener as a handler which receives all
messages sent for the path "a/b" and its sub-paths.

Full usage instructions can be found in the package's online help,
using "go doc github.com/seehuhn/trace".
