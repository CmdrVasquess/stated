package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"

	"git.fractalqb.de/fractalqb/gomk"
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
	}
)

func mkGenerate() {
	mkVersioner()
	build.WDir().Exec("go", "generate", "./...")
}

func mkTest() {
	mkGenerate()
	build.WDir().Exec("go", "test", "./...")
}

func mkBuild() {
	mkGenerate()
	build.WDir().Exec("go", "build", "--trimpath", "./...")
}

func mkVersioner() {
	_, err := exec.LookPath("versioner")
	if err == nil {
		return
	}
	build.WithEnv(func(e *gomk.Env) {
		e.Set("GO111MODULE", "on")
	}, func() {
		build.WDir().Exec("go", "get", "-u", "git.fractalqb.de/fractalqb/pack/versioner")
	})
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
