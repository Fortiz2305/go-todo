package todo

import (
	"fmt"
	"github.com/gonuts/commander"
	"os"
	"time"
)

const (
	todoFilename = ".todo.json"
)

type Task struct {
	Todo   string
	Date   time.Time
	Status string
}

func Run() error {
	command := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "go_todo",
	}

	command.Subcommands = []*commander.Command{
		todoList(todoFilename),
		todoSave(todoFilename),
		todoStatus(todoFilename),
		todoDelete(todoFilename),
	}

	err := command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	return err
}
