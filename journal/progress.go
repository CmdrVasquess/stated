package journal

import "github.com/CmdrVasquess/stated/events"

type progressT string

const ProgressEvent = progressT("Progress")

func (t progressT) New() events.Event { return new(Progress) }
func (t progressT) String() string    { return string(t) }

type Progress struct {
	events.Common
	Combat     int
	Trade      int
	Explore    int
	CQC        int
	Federation int
	Empire     int
}

func (_ *Progress) EventType() events.Type { return ProgressEvent }

func init() {
	events.RegisterType(string(ProgressEvent), ProgressEvent)
}
