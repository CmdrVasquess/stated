package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qblog"
	"github.com/CmdrVasquess/stated"
	"github.com/CmdrVasquess/watched"
)

var (
	log    = qblog.New("edjread")
	logCfg = c4hgol.NewLogGroup(log, "",
		stated.LogCfg,
	)

	edstate = stated.NewEDState(nil)
	ser     int64
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func readJournal(file string) {
	log.Infov("read `journal file`", file)
	rd, err := os.Open(file)
	must(err)
	defer rd.Close()
	scn := bufio.NewScanner(rd)
	for scn.Scan() {
		ser++
		edstate.OnJournalEvent(watched.JounalEvent{
			Serial: ser,
			Event:  scn.Bytes(),
		})
	}
}

func main() {
	fLog := flag.String("log", "", c4hgol.FlagDoc())
	fLsLog := flag.Bool("log-list", false, "List configurable loggers")
	fCmdr := flag.String("cmdr", "", "dir for commander file")
	flag.Parse()
	c4hgol.Configure(logCfg, *fLog, true)
	if *fCmdr != "" {
		edstate.CmdrFile = stated.CmdrFile{Dir: *fCmdr}.Filename
	}
	if *fLsLog {
		wr := flag.CommandLine.Output()
		fmt.Fprintln(wr, "Loggers:")
		c4hgol.Visit(logCfg, func(_ c4hgol.LogConfig, p string) error {
			fmt.Fprintf(wr, " - %s", p)
			return nil
		})
	}
	for _, f := range flag.Args() {
		readJournal(f)
	}
	edstate.Close()
}
