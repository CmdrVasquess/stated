package journal

import "github.com/CmdrVasquess/stated/events"

type materialcollectedT string

const MaterialCollectedEvent = materialcollectedT("MaterialCollected")

func (t materialcollectedT) New() events.Event { return new(MaterialCollected) }
func (t materialcollectedT) String() string    { return string(t) }

type MaterialCollected struct {
	events.Common
	Category string
	Name     string
	Count    int
}

func (_ *MaterialCollected) EventType() events.Type { return MaterialCollectedEvent }

func init() {
	events.RegisterType(string(MaterialCollectedEvent), MaterialCollectedEvent)
}
