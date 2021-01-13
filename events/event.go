package events

import (
	"fmt"
	"sync"
	"time"
)

type Event interface {
	EventType() Type
	Timestamp() time.Time
	Event() string
}

type Common struct {
	Time time.Time `json:"timestamp"`
	Tag  string    `json:"event"`
}

func (c *Common) Timestamp() time.Time { return c.Time }

func (c *Common) Event() string { return c.Tag }

type Type interface {
	fmt.Stringer
	New() Event
}

func MustRegisterType(name string, t Type) {
	err := RegisterType(name, t)
	if err != nil {
		panic(err)
	}
}

func RegisterType(name string, t Type) error {
	if val, ok := eventTypes.LoadOrStore(name, t); ok {
		rt := val.(Type)
		t1, t2 := rt.New(), t.New()
		return fmt.Errorf("duplicate event registration on '%s': %T / %T",
			name,
			t1,
			t2)
	}
	return nil
}

func EventType(event string) Type {
	res, _ := eventTypes.Load(event)
	if res == nil {
		return nil
	}
	return res.(Type)
}

func EventNames() (res []string) {
	eventTypes.Range(func(k, _ interface{}) bool {
		res = append(res, k.(string))
		return true
	})
	return res
}

var eventTypes sync.Map
