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
