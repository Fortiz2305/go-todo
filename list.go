package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"io"
	"os"
	"strings"
)

func todo_list(filename string) *commander.Command {
	list := func(cmd *commander.Command, args []string) error {

		return &commander.Command{
			Run:       cmd_list,
			UsageLine: "list [options]",
			Short:     "show list index",
		}
	}
}
