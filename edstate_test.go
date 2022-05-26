package stated

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/CmdrVasquess/stated/events"
	"github.com/CmdrVasquess/stated/journal"
	"github.com/fractalqb/change"

	"github.com/CmdrVasquess/watched"
)

var _ watched.EventRecv = (*EDState)(nil)

func ExampleEDState_marshal() {
	eds := NewEDState(&Config{
		CmdrFile: CmdrFile{Dir: "."}.Filename,
	})
	eds.Cmdr = NewCommander("F4711")
	eds.Cmdr.Name = change.NewVal("John Doe")
	_, err := json.Marshal(eds)
	fmt.Println(err)
	// Output:
	// <nil>
}

func TestEDState_OnJournalEvent_NoCmdr(t *testing.T) {
	eds := NewEDState(nil)
	fsdjump := journal.FSDJump{ // Must have commander
		Common: events.Common{
			Time: time.Now(),
			Tag:  journal.FSDJumpEvent.String(),
		},
		SSysInfo: journal.SSysInfo{
			SystemAddress: 4711,
			StarSystem:    "Pusemuckel",
			StarPos:       [3]float32{1, 2, 3},
		},
	}
	re, err := json.Marshal(&fsdjump)
	if err != nil {
		t.Fatal(err)
	}
	err = eds.OnJournalEvent(watched.JounalEvent{
		Serial: 1,
		Event:  re,
	})
	if err == nil {
		t.Error("expected error for missing commander")
	}
}
