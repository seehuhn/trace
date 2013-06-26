package trace

import (
	"fmt"
	"time"
)

type Consumer func(t time.Time, module string, component string, msg string)

type ConsumerId int

var consumers map[ConsumerId]Consumer

var consumerIdx ConsumerId

func Register(c Consumer) ConsumerId {
	handle := consumerIdx
	consumerIdx += 1
	consumers[handle] = c
	return handle
}

func Unregister(handle ConsumerId) {
	delete(consumers, handle)
}

func T(module, component, tmpl string, args ...interface{}) {
	if consumers == nil {
		return
	}
	t := time.Now()
	msg := fmt.Sprintf(tmpl, args...)
	for _, c := range consumers {
		c(t, module, component, msg)
	}
}

func init() {
	consumers = map[ConsumerId]Consumer{}
}
