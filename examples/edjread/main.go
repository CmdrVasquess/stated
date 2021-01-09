package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qbsllm"
	"github.com/CmdrVasquess/stated"
	"github.com/CmdrVasquess/watched"
)

var (
	log                      = qbsllm.New(qbsllm.Lnormal, "edjread", nil, nil)
	LogCfg c4hgol.Configurer = c4hgol.Config(qbsllm.NewConfig(log),
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
	log.Infoa("read `journal file`", file)
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
	fLog := flag.String(c4hgol.DefaultFlagLevel, "", c4hgol.LevelCfgDoc(nil))
	fLsLog := flag.Bool(c4hgol.DefaultFlagList, false, "List configurable loggers")
	fCmdr := flag.String("cmdr", "", "dir for commander file")
	flag.Parse()
	c4hgol.SetLevel(LogCfg, *fLog, nil)
	if *fCmdr != "" {
		edstate.CmdrFile = stated.CmdrFile{Dir: *fCmdr}.Filename
	}
	if *fLsLog {
		wr := flag.CommandLine.Output()
		fmt.Fprintln(wr, "Loggers:")
		c4hgol.ListLogs(LogCfg, wr, " - ")
	}
	for _, f := range flag.Args() {
		readJournal(f)
	}
	edstate.Close()
}
