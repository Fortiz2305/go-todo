package todo

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/gonuts/commander"
)

func todoStatus(tasksFile string) *commander.Command {
	setStatus := func(cmd *commander.Command, args []string) error {
		file, err := ioutil.ReadFile(tasksFile)
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(file), "\n")

		for i, line := range lines {
			if strings.Contains(line, "ID\": "+"\""+os.Args[2]) {
				lines[i+3] = "  \"Status\":" + " \"" + os.Args[3] + "\""
			}
		}

		output := strings.Join(lines, "\n")
		errorWriting := ioutil.WriteFile(tasksFile, []byte(output), 0644)
		if errorWriting != nil {
			panic(errorWriting)
		}

		return nil
	}

	return &commander.Command{
		Run:       setStatus,
		UsageLine: "setStatus Task Status",
		Short:     "Set the status of your task",
		Long: `
Set the status of your task.

ex:
  $ ./go-todo setStatus "MyTask" "MyStatus"
`,
	}
}
