// Package trace implements a simple tracing framework which can be
// used to diagnose run-time problems.
//
// Sending Messages
//
// Code using this framework can emit diagnostic messages using the
// trace.T function.  Example:
//
//     trace.T("a/b/c", trace.PrioError,
//             "failed to connect to server %q, using offline mode", serverName)
//
// The first argument in this call is a path which gives information
// about the origin of the message, the second argument indicates the
// importance of the message.  Both, the path and the priority are
// used to decide which listeners receive the correponding message.
//
// Receiving Messages
//
// Listeners can subscribe to messages, either for a given path or for
// all paths, and they can specify a minimum priority for messages to
// be delivered.  Example:
//
//     func MyListener(t time.Time, path string, prio Priority, msg string) {
//             // ... do something
//     }
//
//     func main() {
//             listener := trace.Register(MyListener, "a/b", trace.PrioAll)
//             // ... code which calls trace.T()
//             listener.Unregister()
//     }
//
// This code installs MyListener as a handler which receives all
// messages sent for the path "a/b" and its sub-paths.
package trace
