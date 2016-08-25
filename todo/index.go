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

/*
Task structure
  Todo: Task name
	Date: Task Date
	Status: Task Status
*/
type Task struct {
	Todo   string
	Date   time.Time
	Status string
}

/*
Run function

It launches the go-todo CLI and return an error variable
*/
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
