package journal

import "github.com/CmdrVasquess/stated/events"

type rankT string

const RankEvent = rankT("Rank")

func (t rankT) New() events.Event { return new(Rank) }
func (t rankT) String() string    { return string(t) }

type Rank struct {
	events.Common
	Combat     int
	Trade      int
	Explore    int
	CQC        int
	Federation int
	Empire     int
}

func (_ *Rank) EventType() events.Type { return RankEvent }

func init() {
	events.RegisterType(string(RankEvent), RankEvent)
}
