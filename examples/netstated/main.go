package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qblog"
	"github.com/fractalqb/change"

	"github.com/CmdrVasquess/stated"
	"github.com/CmdrVasquess/watched/edeh/edehnet"
)

var (
	recv    edehnet.Receiver
	edstate = stated.NewEDState(&stated.Config{
		CmdrFile: stated.CmdrFile{Dir: "."}.Filename,
	})

	log    = qblog.New("netstated")
	logCfg = c4hgol.NewLogGroup(log, "",
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
		for c := change.Flags(1); c < stated.ChgEND; c = c << 1 {
			if chg.Change&c == 0 {
				continue
			}
			switch c {
			case stated.ChgGame:
				fmt.Printf(" - game : %s beta: %t %s (%s-%s)\n",
					edstate.EDVersion,
					edstate.Beta,
					edstate.Language,
					edstate.L10n.Lang,
					edstate.L10n.Region,
				)
			case stated.ChgCommander:
				fmt.Print(" - commander: ")
				json.NewEncoder(os.Stdout).Encode(edstate.Cmdr)
			case stated.ChgSystem:
				fmt.Printf(" - system %T: ", edstate.Loc.Location)
				json.NewEncoder(os.Stdout).Encode(edstate.Loc.Location)
			case stated.ChgLocation:
				fmt.Printf(" - location %T: ", edstate.Loc.Location)
				json.NewEncoder(os.Stdout).Encode(edstate.Loc.Location)
			default:
				fmt.Printf(" - chg#%d\n", c)
			}
		}
	}
}

func main() {
	fLog := flag.String("log", "", c4hgol.FlagDoc())
	fLsLog := flag.Bool("log-list", false, "List configurable loggers")
	flag.StringVar(&recv.Listen, "l", "", "TCP listen address")
	flag.Parse()
	c4hgol.Configure(logCfg, *fLog, true)
	if *fLsLog {
		wr := flag.CommandLine.Output()
		fmt.Fprintln(wr, "Loggers:")
		c4hgol.Visit(logCfg, func(_ c4hgol.LogConfig, p string) error {
			fmt.Fprintf(wr, " - %s", p)
			return nil
		})
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
