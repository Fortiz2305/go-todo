package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"os"
	"time"
)

const (
	todo_filename = "/Users/Fortiz/.todo.json"
)

type Task struct {
	Todo string
	Date time.Time
	Status string
}

type Tasks struct {
	Collection []Task
}

func main() {
	command := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "go_todo",
	}

	command.Subcommands = []*commander.Command{
		todo_list(todo_filename),
		todo_save(todo_filename),
	}

	err := command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
