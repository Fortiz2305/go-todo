package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"os"
)

const (
	todo_filename = "~/.todo.json"
)

func main() {
	command := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "go_todo",
	}

	command.Subcommands = []*commander.Command{
		todo_list(todo_filename),
		todo_save(todo_filename),
		todo_solve(todo_filename),
	}

	err := command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
