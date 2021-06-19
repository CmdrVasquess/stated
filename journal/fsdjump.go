package journal

import "github.com/CmdrVasquess/stated/events"

type fsdjumpT string

const FSDJumpEvent = fsdjumpT("FSDJump")

func (t fsdjumpT) New() events.Event { return new(FSDJump) }
func (t fsdjumpT) String() string    { return string(t) }

type FSDJump struct {
	events.Common
	JumpDist float32
	Body     string
	BodyID   int
	SSysInfo
}

func (_ *FSDJump) EventType() events.Type { return FSDJumpEvent }

func init() {
	events.MustRegisterType(string(FSDJumpEvent), FSDJumpEvent)
}
