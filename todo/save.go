package todo

import (
  "fmt"
  "github.com/gonuts/commander"
  "encoding/json"
  "bytes"
  "os"
  "time"
)

func todo_save(tasks_file string) *commander.Command {
  save := func(cmd *commander.Command, args []string) error {
    if len(args) == 0 {
      cmd.Usage()
      return nil
    }

    newTask := Task{os.Args[2], time.Now().Local(), "OPEN"}

    task_json, _ := json.Marshal(newTask)
    task_json, _ = prettyprint(task_json)
    file, err := os.OpenFile(tasks_file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
    if err != nil {
      panic(err)
    }
    defer file.Close()

    _, err = fmt.Fprintf(file, "%s\n", task_json)
    if err != nil {
      panic(err)
    }
    return nil
  }

  return &commander.Command {
    Run: save,
    UsageLine: "save [task]",
    Short: "Save a new task",
    Long:`
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