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

// Package trace implements a simple tracing framework which can be
// used to diagnose run-time problems.
//
// Sending Messages
//
// Code using this framework can emit diagnostic messages using the
// trace.T() function.  Example:
//
//     trace.T("client/setup", trace.PrioError,
//             "failed to connect to server %q, using offline mode", serverName)
//
// The first argument in this call is a path which gives information
// about the origin of the message, the second argument indicates the
// importance of the message.  Both, the path and the priority are
// used to decide which listeners receive the correponding message.
// The following arguments, a format string and additional optional
// arguments, are passed to fmt.Sprintf to compose the message
// reported to the listeners registered for the given message path.
//
// Receiving Messages
//
// Listeners can subscribe to messages, either for a given path or for
// all paths, using the Register() method.  A minimum priority for
// messages to be delivered can be used.  Example:
//
//     func printTrace(t time.Time, path string, prio trace.Priority, msg string) {
//             fmt.Printf("%s:%s: %s\n", t.Format("15:04:05.000"), path, msg)
//     }
//
//     func main() {
//             listener := trace.Register(printTrace, "client", trace.PrioAll)
//             // ... code which calls trace.T()
//             listener.Unregister()
//     }
//
// This code installs printTrace as a handler which receives all
// messages sent for the path "client" and its sub-paths.
package trace
