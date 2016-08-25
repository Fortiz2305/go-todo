package todo

import (
	"github.com/gonuts/commander"
	"io/ioutil"
	"os"
	"strings"
)

func todo_status(tasks_file string) *commander.Command {
	set_status := func(cmd *commander.Command, args []string) error {
		file, err := ioutil.ReadFile(tasks_file)
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(file), "\n")

		for i, line := range lines {
			if strings.Contains(line, os.Args[2]) {
				lines[i+2] = "  \"Status\":" + " \"" + os.Args[3] + "\""
			}
		}

		output := strings.Join(lines, "\n")
		error_writing := ioutil.WriteFile(tasks_file, []byte(output), 0644)
		if error_writing != nil {
			panic(error_writing)
		}

		return nil
	}

	return &commander.Command{
		Run:       set_status,
		UsageLine: "set_status Task Status",
		Short:     "Set the status of your task",
    Long: `
Set the status of your task.

ex:
  $ ./go-todo set_status "MyTask" "MyStatus"
`,
	}
}
