package todo

import (
	"fmt"
	"github.com/gonuts/commander"
	"os"
	"time"
)

const (
	todo_filename = ".todo.json"
)

type Task struct {
	Todo string
	Date time.Time
	Status string
}

func Run() (error) {
	command := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "go_todo",
	}

	command.Subcommands = []*commander.Command{
		todo_list(todo_filename),
		todo_save(todo_filename),
		todo_status(todo_filename),
		todo_delete(todo_filename),
	}

	err := command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	return err
}
