package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gonuts/commander"
)

func todoSave(tasksFile string) *commander.Command {
	save := func(cmd *commander.Command, args []string) error {
		if len(args) == 0 {
			cmd.Usage()
			return nil
		}

		file, err := os.OpenFile(tasksFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		newTask := Task{strconv.Itoa(findID(tasksFile)), os.Args[2], time.Now().Local(), "OPEN"}

		taskJSON, _ := json.Marshal(newTask)
		taskJSON, _ = prettyprint(taskJSON)

		_, err = fmt.Fprintf(file, "%s\n", taskJSON)
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

func findID(filename string) int {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	n := 1
	for _, line := range lines {
		if strings.Contains(line, "ID") {
			if strings.Contains(line, strconv.Itoa(n)) {
				n++
			} else {
				break
			}
		}
	}
	return n
}
