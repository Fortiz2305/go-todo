package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	BANNER = `
Command Line Interface to create tasks TODO
Version: %s
  `
	VERSION = "v0.0.1"
)

var (
	help   bool
	create bool
	list   bool
	solve  bool
)

func init() {
	flag.BoolVar(&help, "help", false, "Help menu")
	flag.BoolVar(&create, "c", false, "Creates and saves a task")
	flag.BoolVar(&list, "l", false, "Lists current tasks")
	flag.BoolVar(&solve, "s", false, "Solves a task")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if help {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	if create {

	} else if list {

	} else if solve {

	}

}
