package stated

import (
	"encoding/json"

	"github.com/CmdrVasquess/stated/att"
)

type Commander struct {
	FID    string
	Name   att.String
	Ranks  Ranks
	ShipID att.Int
	inShip *Ship
}

func NewCommander(fid string) *Commander {
	return &Commander{FID: fid}
}

type Rank struct {
	Level    int `json:"L"`
	Progress int `json:"P"`
}

//go:generate stringer -type RankType
type RankType att.Int16

const (
	Combat RankType = iota
	Trade
	Explore
	CQC
	Federation
	Empire

	RanksNum
)

type Ranks [RanksNum]Rank

func (rs Ranks) MarshalJSON() ([]byte, error) {
	m := make(map[string]Rank)
	for r, s := range rs {
		m[RankType(r).String()] = s
	}
	return json.Marshal(m)
}

func (rs *Ranks) UnmarshalJSON(raw []byte) error {
	m := make(map[string]Rank)
	err := json.Unmarshal(raw, &m)
	if err != nil {
		return err
	}
	for r := range rs {
		rs[r] = m[RankType(r).String()]
	}
	return nil
}
