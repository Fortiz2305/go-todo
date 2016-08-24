package main

import (
  "github.com/gonuts/commander"
  "io/ioutil"
  "os"
  "strings"
)

func todo_delete(tasks_file string) *commander.Command {
  delete := func(cmd *commander.Command, args []string) error {
    file, err := ioutil.ReadFile(tasks_file)
    if err != nil {
      panic(err)
    }

    lines := strings.Split(string(file), "\n")

    for i, line := range lines {
      if strings.Contains(line, os.Args[2]) {
        for k := 0; i == 3; i++ {
          lines[i+k] = ""
        }
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
    Run: delete,
    UsageLine: "delete Task",
    Short: "Delete a task",
    Long:`
Delete a task.

ex:
  $ ./go-todo delete "MyTask"
`,
  }
}
