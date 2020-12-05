package stated

import "time"

type Jump struct {
	SysAddr uint64
	Arrive  time.Time
}

type JumpHist struct {
	MaxLen int
	Jumps  []Jump
	Last   int
}

func (jh *JumpHist) Jump(addr uint64, t time.Time) {
	if len(jh.Jumps) < jh.MaxLen {
		jh.Last = len(jh.Jumps)
		jh.Jumps = append(jh.Jumps, Jump{addr, t})
	} else {
		i := jh.Last + 1
		if i >= len(jh.Jumps) {
			i = 0
		}
		j := &jh.Jumps[i]
		j.SysAddr = addr
		j.Arrive = t
		jh.Last = i
	}
}
