package todo

import (
	"github.com/gonuts/commander"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func todoDelete(tasksFile string) *commander.Command {
	delete := func(cmd *commander.Command, args []string) error {
		file, err := ioutil.ReadFile(tasksFile)
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(file), "\n")

		for i, line := range lines {
			if strings.Contains(line, os.Args[2]) {
				for k := -1; k <= 3; k++ {
					lines[i+k] = ""
				}
			}
		}

		output, _ := removeEmptyLines(strings.Join(lines, "\n"))

		errorWriting := ioutil.WriteFile(tasksFile, []byte(output), 0644)
		if errorWriting != nil {
			panic(errorWriting)
		}

		return nil
	}

	return &commander.Command{
		Run:       delete,
		UsageLine: "delete Task",
		Short:     "Delete a task",
		Long: `
Delete a task.

ex:
  $ ./go-todo delete "MyTask"
`,
	}
}

func removeEmptyLines(lines string) (string, error) {
	regex, err := regexp.Compile("\n\n")
	if err != nil {
		panic(err)
	}
	lines = regex.ReplaceAllString(lines, "\n")

	return lines, err
}
