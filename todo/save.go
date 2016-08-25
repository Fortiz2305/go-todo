package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gonuts/commander"
	"os"
	"time"
)

func todoSave(tasksFile string) *commander.Command {
	save := func(cmd *commander.Command, args []string) error {
		if len(args) == 0 {
			cmd.Usage()
			return nil
		}

		newTask := Task{os.Args[2], time.Now().Local(), "OPEN"}

		task_JSON, _ := json.Marshal(newTask)
		task_JSON, _ = prettyprint(task_JSON)
		file, err := os.OpenFile(tasksFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = fmt.Fprintf(file, "%s\n", task_JSON)
		if err != nil {
			panic(err)
		}
		return nil
	}

	return &commander.Command{
		Run:       save,
		UsageLine: "save [task]",
		Short:     "Save a new task",
		Long: `
Save a new task.

ex:
  $ ./go-todo save "MyTask"
`,
	}
}

func prettyprint(b []byte) ([]byte, error) {
	var output bytes.Buffer
	err := json.Indent(&output, b, "", "  ")
	return output.Bytes(), err
}
