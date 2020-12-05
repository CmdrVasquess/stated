package stated

import (
	"os"
	"path/filepath"

	"github.com/CmdrVasquess/stated/att"
)

type RawMatStats struct {
	Min, Max float32
	Sum      float64
	Count    int
}

type Commander struct {
	FID         string
	Name        att.String
	Ranks       Ranks
	ShipID      att.Int
	Mats        Materials
	inShip      *Ship
	RawMatStats map[string]*RawMatStats
}

func NewCommander(fid string) *Commander {
	return &Commander{
		FID:         fid,
		RawMatStats: make(map[string]*RawMatStats),
	}
}

func (cmdr *Commander) Save(file string) error {
	dir := filepath.Dir(file)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Infoa("create `commander` `dir`", cmdr.Name, dir)
		if err := os.Mkdir(dir, 0777); err != nil {
			log.Errore(err)
		}
	}
	log.Infoa("save `commander` with `fid` to `file`", cmdr.Name, cmdr.FID, file)
	return SaveJSON(file, cmdr, "")
}

type Rank struct {
	Level    int
	Progress int
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
