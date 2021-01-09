package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"git.fractalqb.de/fractalqb/gomk"
	"git.fractalqb.de/fractalqb/gomk/task"
)

const (
	defaultTagret = "test"
)

var (
	build   = gomk.MustNewBuild("", os.Environ())
	targets = map[string]func(){
		"generate":    mkGenerate,
		defaultTagret: mkTest,
		"build":       mkBuild,
		"deps":        func() { task.DepsGraph(build) },
	}
)

func mkGenerate() {
	task.GetVersioner(build)
	build.WDir().Exec("go", "generate", "./...")
}

func mkTest() {
	mkGenerate()
	build.WDir().Exec("go", "test", "./...")
}

func mkBuild() {
	mkGenerate()
	build.WDir().Cd("tools", "newjournalevent").Exec("go", "build", "--trimpath")
	build.WDir().Cd("examples", "netstated").Exec("go", "build", "--trimpath")
}

func usage() {
	wr := flag.CommandLine.Output()
	fmt.Fprintf(wr, "tagrets for module 'stated' (default: %s):\n", defaultTagret)
	var ts []string
	for t, _ := range targets {
		ts = append(ts, t)
	}
	sort.Strings(ts)
	for _, t := range ts {
		fmt.Fprintf(wr, " - %s\n", t)
	}
}

func buildTarget(name string) {
	recipe := targets[name]
	if recipe == nil {
		log.Fatalf("unknown target: '%s'", name)
	}
	err := gomk.Try(recipe)
	if err != nil {
		log.Fatalf("target '%s' failed: %s", name, err)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()
	log.Printf("project root: %s\n", build.PrjRoot)
	if len(flag.Args()) == 0 {
		buildTarget(defaultTagret)
	}
	for _, target := range flag.Args() {
		buildTarget(target)
	}
}
