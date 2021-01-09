package main

import (
	"flag"
	"fmt"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qbsllm"

	"github.com/CmdrVasquess/stated"
	"github.com/CmdrVasquess/stated/att"
	"github.com/CmdrVasquess/watched/edeh/edehnet"
)

var (
	recv    edehnet.Receiver
	edstate = stated.NewEDState(&stated.Config{
		CmdrFile: stated.CmdrFile{Dir: "."}.Filename,
	})

	log                      = qbsllm.New(qbsllm.Lnormal, "netstated", nil, nil)
	LogCfg c4hgol.Configurer = c4hgol.Config(qbsllm.NewConfig(log),
		edehnet.LogCfg,
		stated.LogCfg,
	)

	changes = make(chan stated.ChangeEvent)
)

func printChanges() {
	for chg := range changes {
		fmt.Printf("%s: %s event changed:\n",
			chg.Event.Timestamp(),
			chg.Event.Event())
		for c := att.Change(1); c < stated.ChgEND; c = c << 1 {
			if chg.Change&c == 0 {
				continue
			}
			switch c {
			case stated.ChgGame:
				fmt.Printf(" - game: %s\n", edstate.EDVersion)
			case stated.ChgCommander:
				fmt.Printf(" - commander: %+v\n", edstate.Cmdr)
			case stated.ChgSystem:
				fmt.Printf(" - system: %+v\n", edstate.Loc.Location)
			case stated.ChgLocation:
				fmt.Printf(" - location: %+v\n", edstate.Loc.Location)
			default:
				fmt.Printf(" - chg#%d\n", c)
			}
		}
	}
}

func main() {
	fLog := flag.String(c4hgol.DefaultFlagLevel, "", c4hgol.LevelCfgDoc(nil))
	fLsLog := flag.Bool(c4hgol.DefaultFlagList, false, "List configurable loggers")
	flag.StringVar(&recv.Listen, "l", "", "TCP listen address")
	flag.Parse()
	c4hgol.SetLevel(LogCfg, *fLog, nil)
	if *fLsLog {
		wr := flag.CommandLine.Output()
		fmt.Fprintln(wr, "Loggers:")
		c4hgol.ListLogs(LogCfg, wr, " - ")
	}
	edstate.Notify = []chan<- stated.ChangeEvent{changes}
	go printChanges()
	log.Infof("StatED receiving edeh-net events on %s…", recv.Listen)
	for {
		err := recv.Run(edstate)
		if err != nil {
			log.Errore(err)
			break
		}
	}
}
