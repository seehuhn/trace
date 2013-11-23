package trace

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func flagListener(t time.Time, path string, prio Priority, msg string) {
	fmt.Printf("%s:%s: %s\n", t.Format("15:04:05.000"), path, msg)
}

type traceInfo struct {
	handle ListenerHandle
	prio   Priority
	path   string
}

func (t *traceInfo) Set(value string) error {
	if t != nil {
		t.handle.Unregister()
	}
	t = nil

	parts := strings.SplitN(value, "@", 2)
	var prio Priority
	switch parts[0] {
	case "none":
		return nil
	case "critical":
		prio = PrioCritical
	case "error":
		prio = PrioError
	case "true", "info":
		prio = PrioInfo
	case "debug":
		prio = PrioDebug
	case "verbose":
		prio = PrioVerbose
	case "all":
		prio = PrioAll
	default:
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("cannot parse priority %q", parts[0])
		}
		prio = Priority(x)
	}

	var path string
	if len(parts) > 1 {
		path = parts[1]
	}

	t = &traceInfo{
		handle: Register(flagListener, path, prio),
		prio:   prio,
		path:   path,
	}
	T("trace", PrioInfo, "tracing %q", t.String())

	return nil
}

func (t *traceInfo) String() string {
	if t == nil {
		return "none"
	}
	var s string
	switch t.prio {
	case PrioCritical:
		s = "critical"
	case PrioError:
		s = "error"
	case PrioInfo:
		s = "info"
	case PrioDebug:
		s = "debug"
	case PrioVerbose:
		s = "verbose"
	case PrioAll:
		s = "all"
	default:
		s = strconv.Itoa(int(t.prio))
	}
	if t.path != "" {
		s = s + "@" + t.path
	}
	return s
}

func (t *traceInfo) IsBoolFlag() bool {
	return true
}

var traceFlag *traceInfo

func init() {
	flag.Var(traceFlag, "trace", "enable tracing for priority@path")
}
