package main

import (
	"github.com/gonuts/commander"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"io"
	"os"
	"strings"
	"encoding/json"
)

func todo_list(tasks_file string) *commander.Command {
	list := func(cmd *commander.Command, args []string) error {
		file, err := ioutil.ReadFile(tasks_file)
		if err != nil {
			panic(err)
		}

		dec := json.NewDecoder(strings.NewReader(string(file)))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Task", "Date", "Status"})

		for {
			var t Task
			if err := dec.Decode(&t); err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			data := []string{t.Todo, t.Date.Format("2006-01-12"), t.Status}
			table.Append(data)
		}
		table.Render()

		return nil
	}

	return &commander.Command {
		Run:       list,
		UsageLine: "list",
		Short:     "Show the task list",
		Long:`
Show the task list.

ex:
  $ ./go-todo list
`,
	}
}
